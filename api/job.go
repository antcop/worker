/**
 * Ant Worker Project
 *
 * Copyright (c) 2015 Epinion Online Research Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Jerry Pham       <jerry@andjerry.com>
 *     Loi Nguyen       <loint@penlook.com>
 */

package api

import (
	"github.com/gin-gonic/gin"
)

func createJob(c *gin.Context) {
	c.JSON(200, gin.H {
		"status": "Create new job",
	})
}

func getJobs(c *gin.Context) {
	c.JSON(200, gin.H {
		"status": "Retrieve all jobs",
	})
}

func getJob(c *gin.Context) {
	c.JSON(200, gin.H {
		"status": "Retrieve one job",
	})
}

func updateJob(c *gin.Context) {
	c.JSON(200, gin.H {
		"status": "Update one job",
	})
}

func partlyUpdateJob(c *gin.Context) {
	c.JSON(200, gin.H {
		"status": "Update parly job",
	})
}

func deleteJob(c *gin.Context) {
	c.JSON(200, gin.H {
		"status": "Delete job",
	})
}