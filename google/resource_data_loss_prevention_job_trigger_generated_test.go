// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	triggers {
		schedule {
			recurrence_period_duration = "86400s"
		}
	}

	inspect_job {
		inspect_template_name = "fake"
		actions {
			save_findings {
				output_config {
					table {
						project_id = "project"
						dataset_id = "dataset"
					}
				}
			}
		}
		storage_config {
			cloud_storage_options {
				file_set {
					url = "gs://mybucket/directory/"
				}
			}
		}
	}
}
`, context)
}

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerBigqueryRowLimitExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerBigqueryRowLimitExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.bigquery_row_limit",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerBigqueryRowLimitExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "bigquery_row_limit" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	triggers {
		schedule {
			recurrence_period_duration = "86400s"
		}
	}

	inspect_job {
		inspect_template_name = "fake"
		actions {
			save_findings {
				output_config {
					table {
						project_id = "project"
						dataset_id = "dataset"
					}
				}
			}
		}
		storage_config {
			big_query_options {
				table_reference {
				    project_id = "project"
				    dataset_id = "dataset"
				    table_id = "table_to_scan"
				}

				rows_limit = 1000
				sample_method = "RANDOM_START"
			}
		}
	}
}
`, context)
}

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerBigqueryRowLimitPercentageExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerBigqueryRowLimitPercentageExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.bigquery_row_limit_percentage",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerBigqueryRowLimitPercentageExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "bigquery_row_limit_percentage" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	triggers {
		schedule {
			recurrence_period_duration = "86400s"
		}
	}

	inspect_job {
		inspect_template_name = "fake"
		actions {
			save_findings {
				output_config {
					table {
						project_id = "project"
						dataset_id = "dataset"
					}
				}
			}
		}
		storage_config {
			big_query_options {
				table_reference {
				    project_id = "project"
				    dataset_id = "dataset"
				    table_id = "table_to_scan"
				}

				rows_limit_percent = 50
				sample_method = "RANDOM_START"
			}
		}
	}
}
`, context)
}

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerDataCatalogOutputExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerDataCatalogOutputExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.data_catalog_output",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerDataCatalogOutputExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "data_catalog_output" {
  parent = "projects/%{project}"
  description = "Description"
  display_name = "Displayname"

  triggers {
    schedule {
      recurrence_period_duration = "86400s"
    }
  }

  inspect_job {
    inspect_template_name = "fake"
    actions {
      publish_findings_to_cloud_data_catalog {
      }
    }
    storage_config {
      big_query_options {
        table_reference {
          project_id = "project"
          dataset_id = "dataset"
          table_id = "table_to_scan"
        }
        rows_limit_percent = 50
        sample_method = "RANDOM_START"
      }
    }
  }
}
`, context)
}

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerSccOutputExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerSccOutputExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.scc_output",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerSccOutputExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "scc_output" {
  parent = "projects/%{project}"
  description = "Description"
  display_name = "Displayname"

  triggers {
    schedule {
      recurrence_period_duration = "86400s"
    }
  }

  inspect_job {
    inspect_template_name = "fake"
    actions {
      publish_summary_to_cscc {
      }
    }
    storage_config {
      big_query_options {
        table_reference {
          project_id = "project"
          dataset_id = "dataset"
          table_id = "table_to_scan"
        }
        rows_limit_percent = 50
        sample_method = "RANDOM_START"
      }
    }
  }
}
`, context)
}

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerJobNotificationEmailsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerJobNotificationEmailsExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.job_notification_emails",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerJobNotificationEmailsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "job_notification_emails" {
  parent       = "projects/%{project}"
  description  = "Description for the job_trigger created by terraform"
  display_name = "TerraformDisplayName"
  
  triggers {
    schedule {
      recurrence_period_duration = "86400s"
    }
  }
  
  inspect_job {
    inspect_template_name = "sample-inspect-template"
    actions {
      job_notification_emails {}
    }
    storage_config {
      cloud_storage_options {
        file_set {
          url = "gs://mybucket/directory/"
        }
      }
    }
  }
}
`, context)
}

func TestAccDataLossPreventionJobTrigger_dlpJobTriggerDeidentifyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"name":          "tf_test_" + RandString(t, 10),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJobTrigger_dlpJobTriggerDeidentifyExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job_trigger.deidentify",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionJobTrigger_dlpJobTriggerDeidentifyExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job_trigger" "deidentify" {
  parent       = "projects/%{project}"
  description  = "Description for the job_trigger created by terraform"
  display_name = "TerraformDisplayName"
  
  triggers {
    schedule {
      recurrence_period_duration = "86400s"
    }
  }
  
  inspect_job {
    inspect_template_name = "sample-inspect-template"
    actions {
      deidentify {
        cloud_storage_output    = "gs://samplebucket/dir/"
        file_types_to_transform = ["CSV", "TSV"]
        transformation_details_storage_config {
          table {
            project_id = "%{project}"
            dataset_id = google_bigquery_dataset.default.dataset_id
            table_id   = google_bigquery_table.default.table_id
          }
        }
        transformation_config {
          deidentify_template            = "sample-deidentify-template"
          image_redact_template          = "sample-image-redact-template"
          structured_deidentify_template = "sample-structured-deidentify-template"
        }
      }
    }
    storage_config {
      cloud_storage_options {
        file_set {
          url = "gs://mybucket/directory/"
        }
      }
    }
  }
}
  
resource "google_bigquery_dataset" "default" {
  dataset_id                  = "%{name}"
  friendly_name               = "terraform-test"
  description                 = "Description for the dataset created by terraform"
  location                    = "US"
  default_table_expiration_ms = 3600000
  
  labels = {
    env = "default"
  }
}
  
resource "google_bigquery_table" "default" {
  dataset_id          = google_bigquery_dataset.default.dataset_id
  table_id            = "%{name}"
  deletion_protection = false
  
  time_partitioning {
    type = "DAY"
  }
  
  labels = {
    env = "default"
  }
  
  schema = <<EOF
    [
    {
      "name": "quantity",
      "type": "NUMERIC",
      "mode": "NULLABLE",
      "description": "The quantity"
    },
    {
      "name": "name",
      "type": "STRING",
      "mode": "NULLABLE",
      "description": "Name of the object"
    }
    ]
  EOF
}
`, context)
}

func testAccCheckDataLossPreventionJobTriggerDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_job_trigger" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/jobTriggers/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("DataLossPreventionJobTrigger still exists at %s", url)
			}
		}

		return nil
	}
}
