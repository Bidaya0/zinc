/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zinclabs/zinc/pkg/auth"
	"github.com/zinclabs/zinc/pkg/meta"
)

// @Summary List User
// @Tags  User
// @Produce json
// @Success 200 {object} meta.SearchResponse
// @Router /api/user [get]
func List(c *gin.Context) {
	users, err := auth.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var Hits []meta.Hit
	for _, u := range users {
		// remove salt from response
		u.Salt = ""
		u.Password = ""
		hit := meta.Hit{
			Index:     u.Name,
			Type:      u.Name,
			ID:        u.ID,
			Timestamp: u.UpdatedAt,
			Source:    u,
		}
		Hits = append(Hits, hit)
	}

	resp := meta.SearchResponse{
		Took: 0,
		Hits: meta.Hits{
			Total: meta.Total{
				Value: len(users),
			},
			MaxScore: 0,
			Hits:     Hits,
		},
	}
	c.JSON(http.StatusOK, resp)
}
