package project

func init() {
	content["/domain/po/po.go"] = objectsTemplate()
	content["/domain/po/shcema.json"] = objectsShcema()
}

func objectsTemplate() string {
	return `
	//Package po generated by 'freedom new-project {{.PackagePath}}'
	package po
	`
}

func objectsShcema() string {
	return `[{
		"tableName": "user",
		"primaryKey": "userId",
		"columns:int": ["userId", "age"],
		"columns:int8": ["sex"],
		"columns:string": ["name", "address"],
		"columns:timestamp": ["created", "updated"]
	},
	{
		"tableName": "admin",
		"primaryKey": "id",
		"columns:int": ["id", "age", "role_id"],
		"columns:string": ["name", "address"],
		"columns:timestamp": ["created", "updated"]
	}
]`
}
