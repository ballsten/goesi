/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.6.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetMarketsRegionIdHistory200Ok. */
//easyjson:json
type GetMarketsRegionIdHistory200OkList []GetMarketsRegionIdHistory200Ok

/* 200 ok object */
//easyjson:json
type GetMarketsRegionIdHistory200Ok struct {
	Average    float32 `json:"average,omitempty"`     /* average number */
	Date       string  `json:"date,omitempty"`        /* The date of this historical statistic entry */
	Highest    float32 `json:"highest,omitempty"`     /* highest number */
	Lowest     float32 `json:"lowest,omitempty"`      /* lowest number */
	OrderCount int64   `json:"order_count,omitempty"` /* Total number of orders happened that day */
	Volume     int64   `json:"volume,omitempty"`      /* Total */
}
