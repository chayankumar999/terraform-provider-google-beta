// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package datalossprevention_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataLossPreventionDiscoveryConfig_Update(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"basic":      testAccDataLossPreventionDiscoveryConfig_BasicUpdate,
		"org":        testAccDataLossPreventionDiscoveryConfig_OrgUpdate,
		"actions":    testAccDataLossPreventionDiscoveryConfig_ActionsUpdate,
		"conditions": testAccDataLossPreventionDiscoveryConfig_ConditionsCadenceUpdate,
		"filter":     testAccDataLossPreventionDiscoveryConfig_FilterUpdate,
		"cloud_sql":  testAccDataLossPreventionDiscoveryConfig_CloudSqlUpdate,
	}
	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccDataLossPreventionDiscoveryConfig_BasicUpdate(t *testing.T) {

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStart(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigUpdate(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_OrgUpdate(t *testing.T) {

	context := map[string]interface{}{
		"organization":  envvar.GetTestOrgFromEnv(t),
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgRunning(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPaused(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_ActionsUpdate(t *testing.T) {

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStart(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActions(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActionsSensitivity(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_ConditionsCadenceUpdate(t *testing.T) {

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStart(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigConditionsCadence(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_FilterUpdate(t *testing.T) {

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStart(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigFilterRegexesAndConditions(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_CloudSqlUpdate(t *testing.T) {

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStartCloudSql(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigUpdateCloudSql(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent", "last_run_time", "update_time"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStart(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/%{location}"
	location = "%{location}"
	display_name = "display name"
	status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "custom_type" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		custom_info_types {
			info_type {
				name = "MY_CUSTOM_TYPE"
			}

			likelihood = "UNLIKELY"

			regex {
				pattern = "test*"
			}
		}
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/%{location}"
	location = "%{location}"
	status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
			conditions {
				or_conditions {
					min_row_count = 10
					min_age = "10800s"
				}
			}
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.custom_type.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActions(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_pubsub_topic" "basic" {
	name = "test-topic"
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/%{location}"
	location = "%{location}"
	status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
	actions {
        export_data {
            profile_table {
                project_id = "%{project}"
                dataset_id = "dataset"
                table_id = "table"
            }
        }
    }
    actions { 
        pub_sub_notification {
			topic = "projects/%{project}/topics/${google_pubsub_topic.basic.name}"
			event = "NEW_PROFILE"
			pubsub_condition {
				expressions {
					logical_operator = "OR"
					conditions { 
						minimum_risk_score = "HIGH" 
					}
				}
			}
			detail_of_message = "TABLE_PROFILE"
		}
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActionsSensitivity(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_pubsub_topic" "basic" {
	name = "test-topic"
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/%{location}"
	location = "%{location}"
	status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
	actions {
        export_data {
            profile_table {
                project_id = "project"
                dataset_id = "dataset"
                table_id = "table"
            }
        }
    }
    actions { 
        pub_sub_notification {
			topic = "projects/%{project}/topics/${google_pubsub_topic.basic.name}"
			event = "NEW_PROFILE"
			pubsub_condition {
				expressions {
					logical_operator = "OR"
					conditions { 
						minimum_sensitivity_score = "HIGH" 
					}
				}
			}
			detail_of_message = "TABLE_PROFILE"
		}
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgRunning(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "organizations/%{organization}/locations/%{location}"
	location = "%{location}"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
	org_config {
		project_id = "%{project}"
		location {
			organization_id = "%{organization}"
		}
	}
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
	status = "RUNNING"
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPaused(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "organizations/%{organization}/locations/%{location}"
	location = "%{location}"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
	org_config {
		project_id = "%{project}"
		location {
			folder_id = 123
		}
	}
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
	status = "PAUSED"
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigConditionsCadence(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/%{location}"
	location = "%{location}"
	status = "RUNNING"

	targets {
        big_query_target {
            filter {
                other_tables {}
            }
            conditions {
                type_collection = "BIG_QUERY_COLLECTION_ALL_TYPES"
            }
            cadence {
                schema_modified_cadence {
                    types = ["SCHEMA_NEW_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
                table_modified_cadence {
                    types = ["TABLE_MODIFIED_TIMESTAMP"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigFilterRegexesAndConditions(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Display"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
	}
}

resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/%{location}"
	location = "%{location}"
	status = "RUNNING"

	targets {
        big_query_target {
            filter {
                tables {
                    include_regexes {
                        patterns {
                            project_id_regex = ".*"
                            dataset_id_regex = ".*"
                            table_id_regex = ".*"
                        }
                    }
                }
            }
            conditions {
                created_after = "2023-10-02T15:01:23Z"
                types {
                    types = ["BIG_QUERY_TABLE_TYPE_TABLE", "BIG_QUERY_TABLE_TYPE_EXTERNAL_BIG_LAKE"]
                }
                or_conditions {
                    min_row_count = 10
                    min_age = "21600s"
                }
            }
        }
    }
    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigStartCloudSql(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
    parent = "projects/%{project}"
    description = "Description"
    display_name = "Display"
    inspect_config {
        info_types {
            name = "EMAIL_ADDRESS"
        }
    }
}
resource "google_data_loss_prevention_discovery_config" "basic" {
    parent = "projects/%{project}/locations/%{location}"
    location = "%{location}"
    status = "RUNNING"
    targets {
        cloud_sql_target {
            filter {
                collection {
                    include_regexes {
                        patterns {
                            project_id_regex = ".*"
                            instance_regex = ".*"
                            database_regex = "do-not-scan.*"
                            database_resource_name_regex = ".*"
                        }
                    }
                }
            }
            conditions {
                database_engines = ["MYSQL", "POSTGRES"]
                types = ["DATABASE_RESOURCE_TYPE_ALL_SUPPORTED_TYPES"]
            }
            disabled {}
        }
    }
    targets {
        cloud_sql_target {
            filter {
                others {}
            }
            generation_cadence {
                schema_modified_cadence {
                    types = ["NEW_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_MONTHLY"
                }
                refresh_frequency = "UPDATE_FREQUENCY_MONTHLY"
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigUpdateCloudSql(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_inspect_template" "basic" {
    parent = "projects/%{project}"
    description = "Description"
    display_name = "Display"
    inspect_config {
        info_types {
            name = "EMAIL_ADDRESS"  
        }
    }
}
resource "google_data_loss_prevention_discovery_config" "basic" {
    parent = "projects/%{project}/locations/%{location}"
    location = "%{location}"
    status = "RUNNING"
    targets {
        cloud_sql_target {
            filter {
                collection {
                    include_regexes {
                        patterns {
                            project_id_regex = ".*"
                            instance_regex = ".*"
                            database_regex = ".*"
                            database_resource_name_regex = "mytable.*"
                        }
                    }
                }
            }
            conditions {
                database_engines = ["ALL_SUPPORTED_DATABASE_ENGINES"]
                types = ["DATABASE_RESOURCE_TYPE_TABLE"]
            }
            generation_cadence {
                schema_modified_cadence {
                    types = ["NEW_COLUMNS", "REMOVED_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
                refresh_frequency = "UPDATE_FREQUENCY_MONTHLY"
            }
        }
    }
    targets {
        cloud_sql_target {
            filter {
                others {}
            }
            generation_cadence {
                schema_modified_cadence {
                    types = ["NEW_COLUMNS", "REMOVED_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
                refresh_frequency = "UPDATE_FREQUENCY_DAILY"
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}
`, context)
}
