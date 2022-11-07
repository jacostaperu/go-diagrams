cat 10.17.193.149_virtual.json | jq -r '[ 
        to_entries[]| 
          .["name"]=.key |
          .["description"]=.value.description | 
          .["destination"]= .value.destination|  
          .["pool"]=.value.pool | 
          del(.key,.value)
        ] '