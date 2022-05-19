import * as pulumi from "@pulumi/pulumi";
import * as rootly from "@rootly/pulumi";

new rootly.Severity("sev5", {
	name: "SEV5",
	color: "#FF0000"
})

new rootly.Service("elasticsearch_prod", {
  name: "elasticsearch-prod",
  color: "#800080"
})

new rootly.Functionality("add_items_to_card", {
  name: "Add items to card",
  color: "#FFFFFF"
})
