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

import (
	"time"
)

/* A list of GetCorporationsCorporationIdWalletsDivisionJournal200Ok. */
//easyjson:json
type GetCorporationsCorporationIdWalletsDivisionJournal200OkList []GetCorporationsCorporationIdWalletsDivisionJournal200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdWalletsDivisionJournal200Ok struct {
	Amount          float32                                                     `json:"amount,omitempty"`  /* Transaction amount. Positive when value transferred to the first party. Negative otherwise */
	Balance         float32                                                     `json:"balance,omitempty"` /* Wallet balance after transaction occurred */
	Date            time.Time                                                   `json:"date,omitempty"`    /* Date and time of transaction */
	ExtraInfo       GetCorporationsCorporationIdWalletsDivisionJournalExtraInfo `json:"extra_info,omitempty"`
	FirstPartyId    int32                                                       `json:"first_party_id,omitempty"`    /* first_party_id integer */
	FirstPartyType  string                                                      `json:"first_party_type,omitempty"`  /* first_party_type string */
	Reason          string                                                      `json:"reason,omitempty"`            /* reason string */
	RefId           int64                                                       `json:"ref_id,omitempty"`            /* Unique journal reference ID */
	RefType         string                                                      `json:"ref_type,omitempty"`          /* Transaction type, different type of transaction will populate different fields in `extra_info` */
	SecondPartyId   int32                                                       `json:"second_party_id,omitempty"`   /* second_party_id integer */
	SecondPartyType string                                                      `json:"second_party_type,omitempty"` /* second_party_type string */
	Tax             float32                                                     `json:"tax,omitempty"`               /* Tax amount received for tax related transactions */
	TaxRecieverId   int32                                                       `json:"tax_reciever_id,omitempty"`   /* the corporation ID receiving any tax paid */
}
