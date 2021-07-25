package gen

const apifoxTemplate = `
{
	"apiCollection": [
	  {
		"name": "{CamelSystemName}",
		"items": [
		  {
			"name": "{Description}#{CamelName}",
			"items": [
			  {
				"name": "{Description}详情",
				"api": {
				  "id": "5036881",
				  "method": "get",
				  "path": "/{SystemName}/{Name}",
				  "parameters": {
					"query": [],
					"path": []
				  },
				  "responses": [
					{
					  "id": "3508123",
					  "name": "成功",
					  "code": 200,
					  "jsonSchema": { "$ref": "#/definitions/922387" }
					}
				  ],
				  "requestBody": {
					"type": "application/json",
					"parameters": [],
					"jsonSchema": { "$ref": "#/definitions/920500" }
				  },
				  "tags": [],
				  "cases": [],
				  "mocks": []
				}
			  },
			  {
				"name": "{Description}列表",
				"api": {
				  "id": "5036910",
				  "method": "get",
				  "path": "/{SystemName}/{Name}/list",
				  "parameters": { "query": [] },
				  "responses": [
					{
					  "id": "3508152",
					  "name": "成功",
					  "code": 200,
					  "jsonSchema": { "$ref": "#/definitions/922388" }
					}
				  ],
				  "requestBody": {
					"type": "application/json",
					"parameters": [],
					"jsonSchema": { "$ref": "#/definitions/920501" }
				  },
				  "tags": [],
				  "cases": [],
				  "mocks": []
				}
			  },
			  {
				"name": "新增{Description}",
				"api": {
				  "id": "5036916",
				  "method": "post",
				  "path": "/{SystemName}/{Name}",
				  "parameters": { "query": [] },
				  "responses": [
					{
					  "id": "3508158",
					  "name": "成功",
					  "code": 200,
					  "jsonSchema": { "$ref": "#/definitions/920704" }
					}
				  ],
				  "requestBody": {
					"type": "application/json",
					"parameters": [],
					"jsonSchema": { "$ref": "#/definitions/902166" }
				  },
				  "tags": [],
				  "cases": [],
				  "mocks": []
				}
			  },
			  {
				"name": "修改{Description}",
				"api": {
				  "id": "5096880",
				  "method": "patch",
				  "path": "/{SystemName}/{Name}",
				  "parameters": [],
				  "responses": [
					{
					  "id": "3634478",
					  "name": "成功",
					  "code": 200,
					  "jsonSchema": { "$ref": "#/definitions/920704" }
					}
				  ],
				  "requestBody": {
					"type": "application/json",
					"parameters": [],
					"jsonSchema": { "$ref": "#/definitions/928917" }
				  },
				  "tags": [],
				  "cases": [],
				  "mocks": []
				}
			  },
			  {
				"name": "删除{Description}",
				"api": {
				  "id": "5096882",
				  "method": "delete",
				  "path": "/{SystemName}/{Name}",
				  "parameters": [],
				  "responses": [
					{
					  "id": "3634480",
					  "name": "成功",
					  "code": 200,
					  "jsonSchema": { "$ref": "#/definitions/920704" }
					}
				  ],
				  "requestBody": {
					"type": "application/json",
					"parameters": [],
					"jsonSchema": { "$ref": "#/definitions/928921" }
				  },
				  "tags": [],
				  "cases": [],
				  "mocks": []
				}
			  }
			]
		  }
		]
	  }
	],
	"socketCollection": [],
	"schemaCollection": [
	  {
		"name": "Model",
		"items": [
		  {
			"name": "model.{CamelName}",
			"id": "#/definitions/906753",
			"schema": {
			  "jsonSchema": {
				"type": "object",
				"properties": {
				  "name": { "type": "string" }
				}
			  }
			}
		  }
		]
	  },
	  {
		"name": "{SystemName}",
		"items": [
		  {
			"name": "{Description}#{CamelName}",
			"items": [
			  {
				"name": "{CamelName}PostParam",
				"id": "#/definitions/902166",
				"schema": { "jsonSchema": { "type": "object", "properties": {} } }
			  },
			  {
				"name": "{CamelName}GetParam",
				"id": "#/definitions/920500",
				"schema": {
					"jsonSchema": {
						"type": "object",
						"properties": {
							"id": {
							"type": "integer"
							}
						},
						"required": ["id"]
					}
				}
			  },
			  {
				"name": "{CamelName}ListParam",
				"id": "#/definitions/920501",
				"schema": {
				  "jsonSchema": {
					"allOf": [
						{
							"$ref": "#/definitions/960836"
						},
						{
							"type": "object",
							"properties": {}
						}
					]
				}
				}
			  },
			  {
				"name": "{CamelName}PatchParam",
				"id": "#/definitions/928917",
				"schema": {
					"jsonSchema": {
						"type": "object",
						"properties": {
							"id": {
							"type": "integer"
							}
						},
						"required": ["id"]
					}
				}
			  },
			  {
				"name": "{CamelName}DeleteParam",
				"id": "#/definitions/928921",
				"schema": {
					"jsonSchema": {
						"type": "object",
						"properties": {
							"id": {
							"type": "integer"
							}
						},
						"required": ["id"]
					}
				}
			  },
			  {
				"name": "{CamelName}GetRes",
				"id": "#/definitions/922387",
				"schema": {
					"jsonSchema": {
						"type": "object",
						"properties": {
							"id": {
							"type": "integer"
							}
						},
						"required": ["id"]
					}
				}
			  },
			  {
				"name": "{CamelName}ListRes",
				"id": "#/definitions/922388",
				"schema": {
					"jsonSchema": {
						"type": "object",
						"properties": {
							"id": {
							"type": "integer"
							}
						},
						"required": ["id"]
					}
				}
            }
			]
		  }
		]
	  },
	  {
		"name": "Q",
		"items": [
		  {
			"name": "q.Resp",
			"id": "#/definitions/920704",
			"schema": {
			  "jsonSchema": {
				"type": "object",
				"properties": {
				  "code": { "type": "string", "mock": { "mock": "0" } },
				  "msg": { "type": "string" }
				},
				"required": ["code", "msg"]
			  }
			}
		  },
		  {
			"name": "q.Page",
			"id": "#/definitions/960836",
			"schema": {
			  "jsonSchema": {
				"type": "object",
				"properties": {
				  "current": { "type": "integer", "minimum": 1 },
				  "pageSize": { "type": "integer", "minimum": 1, "maximum": 30 }
				},
				"required": ["current", "pageSize"]
			  }
			}
		  }
		]
	  }
	]
  }`
