# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .cause import *
from .functionality import *
from .incident_role import *
from .incident_type import *
from .provider import *
from .service import *
from .severity import *
from .team import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_rootly.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_rootly.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "rootly",
  "mod": "index/cause",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/cause:Cause": "Cause"
  }
 },
 {
  "pkg": "rootly",
  "mod": "index/functionality",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/functionality:Functionality": "Functionality"
  }
 },
 {
  "pkg": "rootly",
  "mod": "index/incidentRole",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/incidentRole:IncidentRole": "IncidentRole"
  }
 },
 {
  "pkg": "rootly",
  "mod": "index/incidentType",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/incidentType:IncidentType": "IncidentType"
  }
 },
 {
  "pkg": "rootly",
  "mod": "index/service",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/service:Service": "Service"
  }
 },
 {
  "pkg": "rootly",
  "mod": "index/severity",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/severity:Severity": "Severity"
  }
 },
 {
  "pkg": "rootly",
  "mod": "index/team",
  "fqn": "pulumi_rootly",
  "classes": {
   "rootly:index/team:Team": "Team"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "rootly",
  "token": "pulumi:providers:rootly",
  "fqn": "pulumi_rootly",
  "class": "Provider"
 }
]
"""
)
