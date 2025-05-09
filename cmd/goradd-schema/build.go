package main

import (
	"context"
	"github.com/goradd/orm/pkg/config"
	db2 "github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/schema"
)

func build(dbConfigFile, inFile, dbKey string) {
	if databaseConfigs, err := config.OpenConfigFile(dbConfigFile); err != nil {
		panic(err)
	} else if err := config.InitDatastore(databaseConfigs); err != nil {
		panic(err)
	} else {
		for _, c := range databaseConfigs {
			if c["key"].(string) == dbKey {
				setDefaultConfigSettings(c)
				db := db2.GetDatabase(c["key"].(string))
				if db == nil {
					panic("database not found")
				}
				if e, ok := db.(db2.SchemaBuilder); ok {
					s, err := schema.ReadJsonFile(inFile)
					if err != nil {
						panic(err)
					}
					ctx := context.Background()
					err = e.DestroySchema(ctx)
					if err != nil {
						panic(err)
					}
					s.Clean()
					err = e.BuildSchema(ctx, *s)
					if err != nil {
						panic(err)
					}
					return
				}
			}
		}
	}
}
