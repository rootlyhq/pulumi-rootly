import * as pulumi from "@pulumi/pulumi";
import * as rootly from "@rootlyhq/pulumi-rootly";

new rootly.Severity("sev5", {
	name: "SEV5",
	color: "#FF0000"
})
