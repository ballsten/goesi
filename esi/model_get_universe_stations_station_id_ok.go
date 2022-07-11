/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.12
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

/* A list of GetUniverseStationsStationIdOk. */
//easyjson:json
type GetUniverseStationsStationIdOkList []GetUniverseStationsStationIdOk

/* 200 ok object */
//easyjson:json
type GetUniverseStationsStationIdOk struct {
	MaxDockableShipVolume    float32                              `json:"max_dockable_ship_volume,omitempty"` /* max_dockable_ship_volume number */
	Name                     string                               `json:"name,omitempty"`                     /* name string */
	OfficeRentalCost         float32                              `json:"office_rental_cost,omitempty"`       /* office_rental_cost number */
	Owner                    int32                                `json:"owner,omitempty"`                    /* ID of the corporation that controls this station */
	Position                 GetUniverseStationsStationIdPosition `json:"position,omitempty"`
	RaceId                   int32                                `json:"race_id,omitempty"`                    /* race_id integer */
	ReprocessingEfficiency   float32                              `json:"reprocessing_efficiency,omitempty"`    /* reprocessing_efficiency number */
	ReprocessingStationsTake float32                              `json:"reprocessing_stations_take,omitempty"` /* reprocessing_stations_take number */
	Services                 []string                             `json:"services,omitempty"`                   /* services array */
	StationId                int32                                `json:"station_id,omitempty"`                 /* station_id integer */
	SystemId                 int32                                `json:"system_id,omitempty"`                  /* The solar system this station is in */
	TypeId                   int32                                `json:"type_id,omitempty"`                    /* type_id integer */
}
