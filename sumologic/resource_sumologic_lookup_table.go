// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//	This file is automatically generated by Sumo Logic and manual
//	changes will be clobbered when the file is regenerated. Do not submit
//	changes to this file.
//
// ----------------------------------------------------------------------------
package sumologic

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSumologicLookupTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicLookupTableCreate,
		Read:   resourceSumologicLookupTableRead,
		Update: resourceSumologicLookupTableUpdate,
		Delete: resourceSumologicLookupTableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"primary_keys": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The primary key field names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ForceNew: true,
			},

			"parent_folder_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1000),
				Required:     true,
			},

			"size_limit_action": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"StopIncomingMessages", "DeleteOldData"}, false),
				Default:      "StopIncomingMessages",
			},

			"fields": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of fields in the lookup table.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"field_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: false,
						},

						"field_type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validation.StringInSlice([]string{"boolean", "int", "long", "double", "string"}, false),
						},
					},
				},
				ForceNew: true,
			},

			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"csv_file_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSumologicLookupTableCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	if d.Id() == "" {
		lookupTable := resourceToLookupTable(d)
		id, err := c.CreateLookupTable(lookupTable)
		if err != nil {
			return err
		}

		d.SetId(id)
	}

	log.Printf("created lookup: %+v\n", d)
	log.Printf("lookup id: %v\n", d.Id())
	return resourceSumologicLookupTableRead(d, meta)
}

func resourceSumologicLookupTableRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id := d.Id()
	lookupTable, err := c.GetLookupTable(id)
	log.Printf("##DEBUG## read lookup: %+v\n", lookupTable)
	if err != nil {
		return err
	}

	if lookupTable == nil {
		log.Printf("[WARN] LookupTable not found, removing from state: %v - %v", id, err)
		d.SetId("")
		return nil
	}

	d.Set("name", lookupTable.Name)
	if err := d.Set("fields", fieldsToList(lookupTable.Fields)); err != nil {
		return fmt.Errorf("error setting fields for resource %s: %s", d.Id(), err)
	}
	d.Set("ttl", lookupTable.Ttl)
	d.Set("primary_keys", lookupTable.PrimaryKeys)
	d.Set("parent_folder_id", lookupTable.ParentFolderId)
	d.Set("size_limit_action", lookupTable.SizeLimitAction)
	d.Set("description", lookupTable.Description)

	return nil
}

func resourceSumologicLookupTableDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	log.Printf("##DEBUG## resourceSumologicLookupTableDelete: %s", d.Id())
	return c.DeleteLookupTable(d.Id())
}

func resourceSumologicLookupTableUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	lookupTable := resourceToLookupTable(d)
	err := c.UpdateLookupTable(lookupTable)
	if err != nil {
		return err
	}

	return resourceSumologicLookupTableRead(d, meta)
}

func resourceToLookupTable(d *schema.ResourceData) LookupTable {

	fieldsData := d.Get("fields").([]interface{})
	var fields []LookupTableField
	for _, data := range fieldsData {
		fields = append(fields, resourceToLookupTableField([]interface{}{data}))
	}

	primaryKeysData := d.Get("primary_keys").([]interface{})
	var primaryKeys []string
	for _, data := range primaryKeysData {
		primaryKeys = append(primaryKeys, data.(string))
	}

	return LookupTable{
		Name:            d.Get("name").(string),
		ID:              d.Id(),
		Fields:          fields,
		Description:     d.Get("description").(string),
		Ttl:             d.Get("ttl").(int),
		SizeLimitAction: d.Get("size_limit_action").(string),
		PrimaryKeys:     primaryKeys,
		ParentFolderId:  d.Get("parent_folder_id").(string),
		CsvFilePath:     d.Get("csv_file_path").(string),
	}
}

func resourceToLookupTableField(data interface{}) LookupTableField {

	lookupTableFieldSlice := data.([]interface{})
	lookupTableField := LookupTableField{}
	if len(lookupTableFieldSlice) > 0 {
		lookupTableFieldObj := lookupTableFieldSlice[0].(map[string]interface{})
		lookupTableField.FieldName = lookupTableFieldObj["field_name"].(string)
		lookupTableField.FieldType = lookupTableFieldObj["field_type"].(string)
	}

	return lookupTableField
}

func fieldsToList(lookupTableField []LookupTableField) []map[string]interface{} {
	var s []map[string]interface{}

	for _, t := range lookupTableField {
		mapping := map[string]interface{}{
			"field_name": t.FieldName,
			"field_type": t.FieldType,
		}
		s = append(s, mapping)
	}

	return s
}
