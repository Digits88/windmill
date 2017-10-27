/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Windmill project.
 *
 * NethServer is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * NethServer is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with NethServer.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package main

import (
	"github.com/gin-gonic/gin"

	"ronzinante/methods"
)

func main() {

	router := gin.Default()

	sessions := router.Group("/api/sessions")
	{
		sessions.GET("/", methods.GetSessions)
		sessions.GET("/:session_id", methods.GetSession)
		sessions.POST("/", methods.CreateSession)
		sessions.PUT("/:server_id", methods.UpdateSession)
		sessions.DELETE("/:server_id", methods.DeleteSession)
	}

	history := router.Group("/api/histories")
	{
		history.GET("/", methods.GetHistories)
		history.GET("/:server_id", methods.GetHistory)
		history.PUT("/:server_id", methods.UpdateHistory)
	}

	router.Run()

}