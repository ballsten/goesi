/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.5.2
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

package goesiv1

/* A list of GetIndustryFacilities200Ok. */
//easyjson:json
type GetIndustryFacilities200OkList []GetIndustryFacilities200Ok

/* 200 ok object */
//easyjson:json
type GetIndustryFacilities200Ok struct {
	FacilityId    int64   `json:"facility_id,omitempty"`     /* ID of the facility */
	OwnerId       int32   `json:"owner_id,omitempty"`        /* Owner of the facility */
	RegionId      int32   `json:"region_id,omitempty"`       /* Region ID where the facility is */
	SolarSystemId int32   `json:"solar_system_id,omitempty"` /* Solar system ID where the facility is */
	Tax           float32 `json:"tax,omitempty"`             /* Tax imposed by the facility */
	TypeId        int32   `json:"type_id,omitempty"`         /* Type ID of the facility */
}
