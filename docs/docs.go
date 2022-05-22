// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/portal/get_greenbirds": {
            "post": {
                "description": "GetGreenbirds",
                "tags": [
                    "Portal"
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"data\": data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/save_greenbirds": {
            "post": {
                "description": "SaveGreenbird",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "description": "新手上路信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GreenbirdData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"保存成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/comment/create": {
            "post": {
                "description": "Create a comment",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户评论成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/comment/like": {
            "post": {
                "description": "Like a comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "LikeCommentData",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LikeCommentData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"点赞成功\", \"commentlike\": commentlike}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/create": {
            "post": {
                "description": "Create a post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "22",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePostData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"发布成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_all_tags": {
            "post": {
                "description": "Get all tags",
                "tags": [
                    "Post"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取标签成功\", \"tags\":tags}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_post_comments": {
            "post": {
                "description": "Get post comments",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取评论成功\", \"comments\":comments}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_post_tags": {
            "post": {
                "description": "Get post tags",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取标签成功\", \"post_tags\":postTags}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "post": {
                "description": "ShowUserInfo",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"查询成功\", \"data\": user}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"登录成功\",\"data\": user}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password1",
                        "name": "password1",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password2",
                        "name": "password2",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/test": {
            "post": {
                "description": "Test",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "用户名，密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Testtype"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreatePostData": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Greenbird": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.GreenbirdData": {
            "type": "object",
            "properties": {
                "greenBirds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Greenbird"
                    }
                }
            }
        },
        "model.LikeCommentData": {
            "type": "object",
            "properties": {
                "comment_id": {
                    "type": "integer"
                },
                "like_or_dislike": {
                    "type": "boolean"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "v1.Testtype": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
