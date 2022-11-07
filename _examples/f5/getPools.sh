cat 10.17.193.149_pool.json |
  jq  -r  '
      to_entries[] |
        .["pool"]=.key  |
        .["description"]=.value.description |
        .["members"]=[
            .value.members |
            to_entries[]|
             .["name"]=.key|
             .["address"]=.value.address|
             .["state"]=.value.state|
             del(.key,.value)
            ] |
        del(.key,.value)

        
    ' | jq -s '.'
