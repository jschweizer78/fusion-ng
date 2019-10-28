package model

// MessageTemplateDetail a single message from API - Root device
type MessageTemplateDetail struct {
	DistributionListIdsCustomizable bool     `json:"distributionListIdsCustomizable"`
	Subject                         string   `json:"subject"`
	DistributionListIds             []string `json:"distributionListIds"`
	UserIdsCustomizable             bool     `json:"userIdsCustomizable"`
}

// NotificationProfileDetail - Child device
type NotificationProfileDetail struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Settings ProfileSettings `json:"settings"`
}

// ProfileSettings - Child device
type ProfileSettings struct {
	ID        string                   `json:"id"`
	Name      string                   `json:"name"`
	Disabled  bool                     `json:"disabled"`
	Value     interface{}              `json:"value"`
	Extension ProfileSettingsExtension `json:"extension"`
}

// ProfileSettingsExtension - Child device
type ProfileSettingsExtension struct {
	Disabled        bool        `json:"disabled"`
	GlobalDisabled  bool        `json:"globalDisabled"`
	EffectiveConfig interface{} `json:"effectiveConfig"`
}

var template = `
{
	"total": 4,
	"partial": true,
	"previous": null,
	"next": "WyJoYWlsIiwiMzdmMGY1MDAtMWJlMC0xMWU0LTkxODEtM2M5NzBlN2ZmNTYwIl0=",
	"data": [
	  {
		"distributionListIdsCustomizable": true,
		"subject": "Lunch options for this Friday",
		"distributionListIds": ["af94ef60-1d9a-11e4-a054-3c970e7ff560"],
		"userIdsCustomizable": false,
		"notificationProfile": {
		  "id": "7026aa0b-720e-11e7-a5a3-696d40dec0a3",
		  "name": "Default Profile",
		  "settings": "... Object omitted for brevity ..."
		},
		"notificationProfileId": {"id": "7026aa0b-720e-11e7-a5a3-696d40dec0a3"},
		"userIDs": ["685c2400-dd09-11e3-8c49-b8e856327746"],
		"confirmationRequestIdDisplay": "auto",
		"bodyDisplay": "auto",
		"audioDisplay": "auto",
		"name": "Friday Food Choices",
		"imageCustomizable": false,
		"subjectCustomizable": false,
		"confirmationRequestIdCustomizable": false,
		"image": 0,
		"imageDisplay": "auto",
		"audioCustomizable": false,
		"audio": 0,
		"alertTone": "default",
		"alertToneCustomizable": false,
		"alertToneDisplay": "auto",
		"confirmationRequestId": "3fb1ee80-1bdf-11e4-9181-3c970e7ff560",
		"createdAt": "2014-08-04T13:58:37.777+0000",
		"expiration": null,
		"expirationCustomizable": false,
		"expirationDisplay": "auto",
		"permissions": ["delete", "put", "get"],
		"distributionListIdsDisplay": "auto",
		"userIdsDisplay": "auto",
		"collaborationGroupIdsDisplay": "auto",
		"collaborationGroupIdsCustomizable": false,
		"collaborationGroupIds": [],
		"deviceGroupIdsDisplay": "auto",
		"deviceGroupIdsCustomizable": false,
		"deviceGroupIds": [],
		"deviceGroups": [],
		"collaborationGroups": [],
		"areaOfInterestIds": null,
		"areaOfInterestIdsCustomizable": false,
		"areasOfInterest": [],
		"areaOfInterestIdsDisplay": "auto",
		"subjectDisplay": "auto",
		"optOut": false,
		"body": "Please take a second to select what you would like to eat this Friday. If you have any dietary restrictions, please let myself or a member of the culinary staff know.",
		"id": "6ec13410-1bdf-11e4-9181-3c970e7ff560",
		"distributionLists": [
		  {
			"createdAt": "2014-08-06T18:51:33.590+0000",
			"id": "af94ef60-1d9a-11e4-a054-3c970e7ff560",
			"name": "Staff"
		  }
		],
		"users": [
		  {
			"createdAt": "2014-05-16T14:50:22.656+0000",
			"id": "685c2400-dd09-11e3-8c49-b8e856327746",
			"name": "Craig Smith"
		  }
		],
		"confirmationRequest": {
		  "escalationRules": [ ],
		  "createdAt": "2014-08-04T13:57:18.824+0000",
		  "id": "3fb1ee80-1bdf-11e4-9181-3c970e7ff560",
		  "options": ["Pizza", "Salad & Bread Sticks"],
		  "name": "Food Options - Friday"
		},
		"bodyCustomizable": false,
		"metadataCustomizable": false,
		"followUp": false
	  }
	]
  }`
