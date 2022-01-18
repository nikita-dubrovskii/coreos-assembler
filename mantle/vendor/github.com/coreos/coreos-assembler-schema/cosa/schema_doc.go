// Generated by ./generate-schema.sh
// DO NOT EDIT

package cosa

var generatedSchemaJSON = `{
 "definitions": {
     "artifact": {
         "type": "object",
         "properties": {
           "path": {
             "$id": "#/artifact/Path",
             "type":"string",
             "title":"Path"
            },
           "sha256": {
             "$id": "#/artifact/sha256",
             "type":"string",
             "title":"SHA256"
            },
           "size": {
             "$id": "#/artifact/size",
             "type":"number",
             "title":"Size in bytes"
            },
           "skip-compression": {
             "$id": "#/artifact/skip-compression",
             "type":"boolean",
             "title":"Skip compression",
             "description":"Artifact should not be compressed or decompressed before use",
             "default":false
            },
           "uncompressed-sha256": {
             "$id": "#/artifact/uncompressed-sha256",
             "type":"string",
             "title":"Uncompressed SHA256"
            },
           "uncompressed-size": {
             "$id": "#/artifact/uncompressed-size",
             "type":"integer",
             "title":"Uncompressed-size"
            }
          },
          "optional": [
              "size",
              "uncompressed-sha256",
              "uncompressed-size",
              "skip-compression"
          ],
          "required": [
              "path",
              "sha256"
          ]
        },
     "image": {
         "type": "object",
         "required": [
             "digest",
             "image"
         ],
         "optional": [
             "comment"
         ],
         "properties": {
           "digest": {
             "$id": "#/image/digest",
             "type":"string",
             "title":"Digest"
            },
           "comment": {
             "$id": "#/image/comment",
             "type":"string",
             "title":"Comment"
            },
           "image": {
             "$id": "#/image/image",
             "type":"string",
             "title":"Image"
            }
          }
      },
      "cloudartifact": {
         "type": "object",
         "required": [
             "url"
         ],
         "optional": [
             "image",
             "object",
             "bucket",
             "region"
         ],
         "properties": {
           "image": {
             "$id":"#/cloudartifact/image",
             "type":"string",
             "title":"Image"
            },
           "url": {
             "$id":"#/cloudartifact/url",
             "type":"string",
             "title":"URL"
            },
            "bucket": {
             "$id":"#/cloudartifact/bucket",
             "type":"string",
             "title":"Bucket"
            },
            "region": {
             "$id":"#/cloudartifact/region",
             "type":"string",
             "title":"Region"
            },
            "object": {
             "$id":"#/cloudartifact/object",
             "type":"string",
             "title":"Object"
            }
          }
     },
     "git": {
         "type": "object",
         "required": [
             "commit",
             "origin"
         ],
         "optional": [
             "branch",
             "dirty"
         ],
         "properties": {
           "branch": {
             "$id":"#/git/branch",
             "type":"string",
             "title":"branch",
             "default":"",
             "examples": [
               "HEAD"
              ],
             "minLength": 3
            },
           "commit": {
             "$id":"#/git/commit",
             "type":"string",
             "title":"commit",
             "default":"",
             "examples": [
               "742edc307e58f35824d906958b6493510e12b593"
              ],
             "minLength": 5
           },
           "dirty": {
             "$id":"#/git/dirty",
             "type":"string",
             "title":"dirty",
             "default":"",
             "examples": [
               "true"
              ],
             "minLength": 1
           },
           "origin": {
             "$id":"#/git/origin",
             "type":"string",
             "title":"origin",
             "default":"",
             "examples": [
               "https://github.com/coreos/fedora-coreos-config"
              ],
             "minLength": 1
            }
          }
     },
     "pkg-items": {
       "type":"array",
       "title":"Package Set differences",
       "items": {
         "$id":"#/pkgdiff/items/item",
         "title":"Items",
         "default":"",
         "minLength": 1
        }
      },
     "advisory-items": {
       "type":"array",
       "title":"Advisory diff",
       "items": {
         "$id":"#/advisory-diff/items/item",
         "title":"Items",
         "default":""
        }
      }
 },
 "$schema":"http://json-schema.org/draft-07/schema#",
 "$id":"http://github.com/coreos/coreos-assembler/blob/main/v1.json.json",
 "type":"object",
 "title":"CoreOS Assember v1 meta.json schema",
 "required": [
     "buildid",
     "name",
     "ostree-commit",
     "ostree-content-checksum",
     "ostree-timestamp",
     "ostree-version",
     "rpm-ostree-inputhash",
     "summary"
 ],
 "optional": [
   "aliyun",
   "amis",
   "azure",
   "azurestack",
   "build-url",
   "digitalocean",
   "exoscale",
   "gcp",
   "ibmcloud",
   "powervs",
   "images",
   "koji",
   "oscontainer",
   "extensions",
   "parent-pkgdiff",
   "pkgdiff",
   "parent-advisories-diff",
   "advisories-diff",
   "release-payload",

   "coreos-assembler.basearch",
   "coreos-assembler.build-timestamp",
   "coreos-assembler.code-source",
   "coreos-assembler.config-dirty",
   "coreos-assembler.config-gitrev",
   "coreos-assembler.container-config-git",
   "coreos-assembler.container-image-git",
   "coreos-assembler.delayed-meta-merge",
   "coreos-assembler.image-config-checksum",
   "coreos-assembler.image-genver",
   "coreos-assembler.image-input-checksum",
   "coreos-assembler.meta-stamp",
   "coreos-assembler.overrides-active",
   "fedora-coreos.parent-commit",
   "fedora-coreos.parent-version",
   "ref"
 ],
 "additionalProperties":false,
 "properties": {
   "ref": {
     "$id":"#/properties/ref",
     "type":"string",
     "title":"BuildRef",
     "default":"",
     "minLength": 1
    },
   "build-url": {
     "$id":"#/properties/build-url",
     "type":"string",
     "title":"Build URL",
     "default":"",
     "minLength": 1
    },
   "buildid": {
     "$id":"#/properties/buildid",
     "type":"string",
     "title":"BuildID",
     "default":"",
     "minLength": 1
    },
   "koji": {
     "type": "object",
         "properties": {
           "build_id": {
             "$id":"#/properties/kojiid",
             "type":"number",
             "title":"Koji Build ID"
           },
           "token": {
             "$id":"#/properties/kojitoken",
             "type":"string",
             "title":"Koji Token"
           },
           "release": {
             "$id":"#/properties/buildrelease",
             "type":"string",
             "title":"Build Release"
           }
         }
   },
   "coreos-assembler.basearch": {
     "$id":"#/properties/coreos-assembler.basearch",
     "type":"string",
     "title":"Architecture",
     "default":"",
     "minLength": 1
    },
   "coreos-assembler.build-timestamp": {
     "$id":"#/properties/coreos-assembler.build-timestamp",
     "type":"string",
     "title":"Build Time Stamp",
     "default":"",
     "minLength": 1
    },
   "coreos-assembler.code-source": {
     "$id":"#/properties/coreos-assembler.code-source",
     "type":"string",
     "title":"CoreOS Source",
     "default":"",
     "minLength": 1
    },
   "coreos-assembler.config-dirty": {
     "$id":"#/properties/coreos-assembler.config-dirty",
     "type":"string",
     "title":"GitDirty",
     "default":"",
     "minLength": 1
    },
   "coreos-assembler.config-gitrev": {
     "$id":"#/properties/coreos-assembler.config-gitrev",
     "type":"string",
     "title":"Config GitRev",
     "default":"",
     "minLength": 1
    },
   "coreos-assembler.container-config-git": {
     "$id":"#/properties/coreos-assembler.container-config-git",
     "type":"object",
     "title":"Container Config GIT",
     "$ref": "#/definitions/git"
    },
   "coreos-assembler.container-image-git": {
     "$id":"#/properties/coreos-assembler.container-image-git",
     "type":"object",
     "title":"COSA Container Image Git",
     "$ref": "#/definitions/git"
    },
   "coreos-assembler.delayed-meta-merge": {
     "$id":"#/properties/coreos-assembler.delayed-meta-merge",
     "type":"boolean",
     "title":"COSA Delayed Meta Merge",
     "default": "False"
    },
   "coreos-assembler.meta-stamp": {
     "$id":"#/properties/coreos-assembler.meta-stamp",
     "type":"number",
     "title":"Meta Stamp",
     "default":"",
     "minLength": 16
    },
    "fedora-coreos.parent-version": {
     "$id":"#/properties/fedora-coreos.parent-version",
     "type":"string",
     "title":"Fedora CoreOS Parent Version",
     "default":"",
     "minLength": 12
    },
    "fedora-coreos.parent-commit": {
     "$id":"#/properties/fedora-coreos.parent-commit",
     "type":"string",
     "title":"Fedora CoreOS parent commit",
     "default":"",
     "examples": [
       "f15f5b25cf138a7683e3d200c53ece2091bf71d31332135da87892ab72ff4ee3"
      ],
     "minLength": 64
    },
   "coreos-assembler.image-config-checksum": {
     "$id":"#/properties/coreos-assembler.image-config-checksum",
     "type":"string",
     "title":"COSA image checksum",
     "default":"",
     "minLength": 64
    },
   "coreos-assembler.image-genver": {
     "$id":"#/properties/coreos-assembler.image-genver",
     "type":"integer",
     "title":"COSA Image Version",
     "default": 0,
     "examples": [
        0
      ]
    },
   "coreos-assembler.image-input-checksum": {
     "$id":"#/properties/coreos-assembler.image-input-checksum",
     "type":"string",
     "title":"Image input checksum",
     "default":"",
     "minLength": 64
    },
   "coreos-assembler.overrides-active": {
     "$id":"#/properties/coreos-assembler.overrides-active",
     "title":"Overrides Active",
     "default":"",
     "type": "boolean"
    },
   "images": {
     "$id":"#/properties/images",
     "type":"object",
     "title":"Build Artifacts",
     "required": [
       "ostree"
      ],
     "optional": [
       "aliyun",
       "aws",
       "azure",
       "azurestack",
       "dasd",
       "digitalocean",
       "exoscale",
       "gcp",
       "ibmcloud",
       "powervs",
       "initramfs",
       "iso",
       "kernel",
       "live-kernel",
       "live-initramfs",
       "live-iso",
       "live-rootfs",
       "metal",
       "metal4k",
       "nutanix",
       "openstack",
       "qemu",
       "vmware",
       "vultr"
      ],
     "properties": {
       "ostree": {
         "$id":"#/properties/images/properties/ostree",
         "type":"object",
         "title":"OSTree",
         "$ref": "#/definitions/artifact"
        },
       "dasd": {
         "$id":"#/properties/images/properties/dasd",
         "type":"object",
         "title":"dasd",
         "$ref": "#/definitions/artifact"
       },
       "exoscale": {
         "$id":"#/properties/images/properties/exoscale",
         "type":"object",
         "title":"exoscale",
         "$ref": "#/definitions/artifact"
       },
       "qemu": {
         "$id":"#/properties/images/properties/qemu",
         "type":"object",
         "title":"Qemu",
         "$ref": "#/definitions/artifact"
        },
       "metal": {
         "$id":"#/properties/images/properties/metal",
         "type":"object",
         "title":"Metal",
         "$ref": "#/definitions/artifact"
        },
       "metal4k": {
         "$id":"#/properties/images/properties/metal4k",
         "type":"object",
         "title":"Metal (4K native)",
         "$ref": "#/definitions/artifact"
        },
       "iso": {
         "$id":"#/properties/images/properties/iso",
         "type":"object",
         "title":"ISO",
         "$ref": "#/definitions/artifact"
        },
       "kernel": {
         "$id":"#/properties/images/properties/kernel",
         "type":"object",
         "title":"Kernel",
         "$ref": "#/definitions/artifact"
        },
       "initramfs": {
         "$id":"#/properties/images/properties/initramfs",
         "type":"object",
         "title":"Initramfs",
         "$ref": "#/definitions/artifact"
        },
       "live-kernel": {
         "$id":"#/properties/images/properties/live-kernel",
         "type":"object",
         "title":"Live Kernel",
         "$ref": "#/definitions/artifact"
         },
       "live-initramfs": {
         "$id":"#/properties/images/properties/live-initramfs",
         "type":"object",
         "title":"Live Initramfs",
         "$ref": "#/definitions/artifact"
        },
       "live-iso": {
        "$id":"#/properties/images/properties/live-iso",
        "type":"object",
        "title":"Live ISO",
        "$ref": "#/definitions/artifact"
       },
       "live-rootfs": {
         "$id":"#/properties/images/properties/live-rootfs",
         "type":"object",
         "title":"Live Rootfs",
         "$ref": "#/definitions/artifact"
        },
        "nutanix": {
          "$id":"#/properties/images/properties/nutanix",
          "type":"object",
          "title":"Nutanix",
          "$ref": "#/definitions/artifact"
         }, 
       "openstack": {
         "$id":"#/properties/images/properties/openstack",
         "type":"object",
         "title":"OpenStack",
         "$ref": "#/definitions/artifact"
        },
       "vmware": {
         "$id":"#/properties/images/properties/vmware",
         "type":"object",
         "title":"VMWare",
         "$ref": "#/definitions/artifact"
        },
       "vultr": {
         "$id": "#/properties/images/properties/vultr",
         "type": "object",
         "title": "Vultr",
         "$ref": "#/definitions/artifact"
        },
       "aliyun": {
         "$id":"#/properties/images/properties/aliyun",
         "type":"object",
         "title":"Aliyun",
         "$ref": "#/definitions/artifact"
        },
       "aws": {
         "$id":"#/properties/images/properties/aws",
         "type":"object",
         "title":"AWS",
         "$ref": "#/definitions/artifact"
       },
       "azure": {
         "$id":"#/properties/images/properties/azure",
         "type":"object",
         "title":"Azure",
         "$ref": "#/definitions/artifact"
        },
       "azurestack": {
         "$id":"#/properties/images/properties/azurestack",
         "type":"object",
         "title":"AzureStack",
         "$ref": "#/definitions/artifact"
       },
       "digitalocean": {
         "$id":"#/properties/images/properties/digitalocean",
         "type":"object",
         "title":"DigitalOcean",
         "$ref": "#/definitions/artifact"
        },
       "ibmcloud": {
         "$id":"#/properties/images/properties/ibmcloud",
         "type":"object",
         "title":"IBM Cloud",
         "$ref": "#/definitions/artifact"
       },
       "powervs": {
        "$id":"#/properties/images/properties/powervs",
        "type":"object",
        "title":"Power Virtual Server",
        "$ref": "#/definitions/artifact"
      },
       "gcp": {
         "$id":"#/properties/images/properties/gcp",
         "type":"object",
         "title":"GCP",
         "$ref": "#/definitions/artifact"
        }
      }
    },
   "name": {
     "$id":"#/properties/name",
     "type":"string",
     "title":"Name",
     "default":"fedora-coreos",
     "examples": [
       "rhcos",
       "fedora-coreos"
      ]
    },
   "oscontainer": {
     "$id":"#/properties/oscontainer",
     "type":"object",
     "title":"Oscontainer",
     "$ref": "#/definitions/image"
    },
   "extensions": {
     "$id":"#/properties/extensions",
     "type":"object",
     "title":"Extensions",
     "required": [
         "path",
         "sha256",
         "rpm-ostree-state",
         "manifest"
     ],
     "properties": {
       "path": {
         "$id": "#/artifact/Path",
         "type":"string",
         "title":"Path"
        },
       "sha256": {
         "$id": "#/artifact/sha256",
         "type":"string",
         "title":"SHA256"
        },
       "rpm-ostree-state": {
         "$id":"#/properties/extensions/items/properties/rpm-ostree-state",
         "type":"string",
         "title":"RpmOstreeState",
         "default":"",
         "minLength": 64
        },
       "manifest": {
         "$id":"#/properties/extensions/items/properties/manifest",
         "type":"object",
         "title":"Manifest"
       }
      }
    },
   "ostree-commit": {
     "$id":"#/properties/ostree-commit",
     "type":"string",
     "title":"ostree-commit",
     "default":"",
     "minLength": 64
    },
   "ostree-content-bytes-written": {
     "$id":"#/properties/ostree-content-bytes-written",
     "type":"integer",
     "title":"ostree-content-bytes-written",
     "default": 0
    },
   "ostree-content-checksum": {
     "$id":"#/properties/ostree-content-checksum",
     "type":"string",
     "title":"ostree-content-checksum",
     "default":"",
     "minLength": 64
    },
   "ostree-n-cache-hits": {
     "$id":"#/properties/ostree-n-cache-hits",
     "type":"integer",
     "title":"ostree-n-cache-hits",
     "default": 0
    },
   "ostree-n-content-total": {
     "$id":"#/properties/ostree-n-content-total",
     "type":"integer",
     "title":"ostree-n-content-total",
     "default": 0
    },
   "ostree-n-content-written": {
     "$id":"#/properties/ostree-n-content-written",
     "type":"integer",
     "title":"ostree-n-content-written",
     "default": 0
    },
   "ostree-n-metadata-total": {
     "$id":"#/properties/ostree-n-metadata-total",
     "type":"integer",
     "title":"ostree-n-metadata-total",
     "default": 0
    },
   "ostree-n-metadata-written": {
     "$id":"#/properties/ostree-n-metadata-written",
     "type":"integer",
     "title":"ostree-n-metadata-written",
     "default": 0
    },
   "ostree-timestamp": {
     "$id":"#/properties/ostree-timestamp",
     "type":"string",
     "title":"ostree timestamp",
     "default":"",
     "examples": [
       "2020-01-15T19:31:31Z"
      ],
     "pattern":"\\d{4}-\\d{2}-\\d{2}T.*Z$"
    },
   "ostree-version": {
     "$id":"#/properties/ostree-version",
     "type":"string",
     "title":"ostree version",
     "default":"",
     "minLength": 1
    },
   "pkgdiff": {
     "$id":"#/properties/pkgdiff",
     "type":"array",
     "title":"pkgdiff between builds",
     "$ref": "#/definitions/pkg-items"
    },
   "parent-pkgdiff": {
     "$id":"#/properties/parent-pkgdiff",
     "type":"array",
     "title":"pkgdiff against parent",
     "$ref": "#/definitions/pkg-items"
    },
   "advisories-diff": {
     "$id":"#/properties/advisories-diff",
     "type":"array",
     "title":"advisory diff between builds",
     "$ref": "#/definitions/advisory-items"
    },
   "parent-advisories-diff": {
     "$id":"#/properties/parent-advisory-diff",
     "type":"array",
     "title":"advisory diff against parent",
     "$ref": "#/definitions/advisory-items"
    },
   "rpm-ostree-inputhash": {
     "$id":"#/properties/rpm-ostree-inputhash",
     "type":"string",
     "title":"input has of the rpm-ostree",
     "default":"",
     "minLength": 64
    },
   "summary": {
     "$id":"#/properties/summary",
     "type":"string",
     "title":"Build Summary",
     "default":"",
     "minLength": 1
    },
   "aliyun": {
     "$id":"#/properties/aliyun",
     "type":"array",
     "title":"Alibaba/Aliyun Uploads",
     "items": {
       "$id":"#/properties/aliyun/images",
       "type":"object",
       "title":"Aliyun Image",
       "required": [
         "name",
         "id"
        ],
       "properties": {
         "name": {
           "$id":"#/properties/aliyun/items/properties/name",
           "type":"string",
           "title":"Region",
           "default":"",
           "minLength": 1
          },
         "id": {
           "$id":"#/properties/aliyun/items/properties/id",
           "type":"string",
           "title":"ImageID",
           "default":"",
           "minLength": 1
          }
        }
      }
    },
   "amis": {
     "$id":"#/properties/amis",
     "type":"array",
     "title":"AMIS",
     "items": {
       "$id":"#/properties/amis/items",
       "type":"object",
       "title":"AMIS",
       "required": [
         "name",
         "hvm",
         "snapshot"
        ],
       "properties": {
         "name": {
           "$id":"#/properties/amis/items/properties/name",
           "type":"string",
           "title":"Region",
           "default":""
          },
         "hvm": {
           "$id":"#/properties/amis/items/properties/hvm",
           "type":"string",
           "title":"HVM",
           "default":""
         },
         "snapshot": {
           "$id":"#/properties/amis/items/properties/snapshot",
           "type":"string",
           "title":"Snapshot",
           "default":""
          }
        }
      }
    },
   "azure": {
     "$id":"#/properties/azure",
     "type":"object",
     "title":"Azure",
     "$ref": "#/definitions/cloudartifact"
    },
   "gcp": {
     "$id":"#/properties/gcp",
     "type":"object",
     "title":"GCP",
     "required": [
         "image",
         "url"
     ],
     "optional": [
         "family",
         "project"
     ],
     "properties": {
       "image": {
         "$id":"#/properties/gcp/image",
         "type":"string",
         "title":"Image Name"
        },
       "url": {
         "$id":"#/properties/gcp/url",
         "type":"string",
         "title":"URL"
        },
       "project": {
         "$id":"#/properties/gcp/project",
         "type":"string",
         "title":"Image Project"
        },
       "family": {
         "$id":"#/properties/gcp/family",
         "type":"string",
         "title":"Image Family"
        }
      }
    },
   "ibmcloud": {
     "$id":"#/properties/ibmcloud",
     "type":"array",
     "title":"IBM Cloud",
     "items": {
       "type":"object",
       "$ref": "#/definitions/cloudartifact"
      }
    },
   "powervs": {
     "$id":"#/properties/powervs",
     "type":"array",
     "title":"Power Virtual Server",
     "items": {
       "type":"object",
       "$ref": "#/definitions/cloudartifact"
      }
    },
    "release-payload": {
      "$id":"#/properties/release-payload",
      "type":"object",
      "title":"ReleasePayload",
      "$ref": "#/definitions/image"
    }
  }
}
`
