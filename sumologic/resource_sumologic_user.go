// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSumologicUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicUserCreate,
		Read:   resourceSumologicUserRead,
		Update: resourceSumologicUserUpdate,
		Delete: resourceSumologicUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"transfer_to": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSumologicUserRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id := d.Id()
	user, err := c.GetUser(id)

	if err != nil {
		return err
	}

	if user == nil {
		log.Printf("[WARN] User not found, removing from state: %v - %v", id, err)
		d.SetId("")
		return nil
	}

	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
	d.Set("email", user.Email)
	if err := d.Set("role_ids", user.RoleIds); err != nil {
		return fmt.Errorf("error setting role ids for resource %s: %s", d.Id(), err)
	}
	d.Set("is_active", user.IsActive)

	return nil
}

func resourceSumologicUserDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)
	return c.DeleteUser(d.Id(), d.Get("transfer_to").(string))
}

func resourceSumologicUserCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	if d.Id() == "" {
		user := resourceToUser(d)
		id, err := c.CreateUser(user)

		if err != nil {
			return err
		}

		d.SetId(id)
	}

	return resourceSumologicUserRead(d, meta)
}

func resourceSumologicUserUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	user := resourceToUser(d)

	err := c.UpdateUser(user)

	if err != nil {
		return err
	}

	return resourceSumologicUserRead(d, meta)
}

func resourceToUser(d *schema.ResourceData) User {
	rawRoleIds := d.Get("role_ids").([]interface{})
	roleIds := make([]string, len(rawRoleIds))
	for i, v := range rawRoleIds {
		roleIds[i] = v.(string)
	}

	return User{
		ID:        d.Id(),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Email:     d.Get("email").(string),
		RoleIds:   roleIds,
		IsActive:  d.Get("is_active").(bool),
	}
}
