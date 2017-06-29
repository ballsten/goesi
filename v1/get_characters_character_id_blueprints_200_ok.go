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

/* A list of GetCharactersCharacterIdBlueprints200Ok. */
//easyjson:json
type GetCharactersCharacterIdBlueprints200OkList []GetCharactersCharacterIdBlueprints200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdBlueprints200Ok struct {
	ItemId             int64  `json:"item_id,omitempty"`             /* Unique ID for this item. The ID of an item is stable if that item is not repackaged, stacked, detached from a stack, assembled, or otherwise altered. If an item is changed in one of these ways, then the ID will also change (see notes below). */
	LocationFlag       string `json:"location_flag,omitempty"`       /* Indicates something about this item's storage location. The flag is used to differentiate between hangar divisions, drone bay, fitting location, and similar. */
	LocationId         int64  `json:"location_id,omitempty"`         /* References a solar system, station or itemID if this blueprint is located within a container. If an itemID the Character - AssetList API must be queried to find the container using the itemID, from which the correct location of the Blueprint can be derived. */
	MaterialEfficiency int32  `json:"material_efficiency,omitempty"` /* Material Efficiency Level of the blueprint, can be any integer between 0 and 10. */
	Quantity           int32  `json:"quantity,omitempty"`            /* Typically will be -1 or -2 designating a singleton item, where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (no activities performed on them yet). */
	Runs               int32  `json:"runs,omitempty"`                /* Number of runs remaining if the blueprint is a copy, -1 if it is an original. */
	TimeEfficiency     int32  `json:"time_efficiency,omitempty"`     /* Time Efficiency Level of the blueprint, can be any even integer between 0 and 20. */
	TypeId             int32  `json:"type_id,omitempty"`             /* type_id integer */
}
