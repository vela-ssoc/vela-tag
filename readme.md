# 标签
> 自动节点标签

```lua
    local tag = vela.tag()
    local host = vela.host

    tag.add(runtime.ARCH)
    
    
    if runtime.OS == "linux" then
        local version = host.version
        tag.add(version:sub(1,1) ..".x")
    end

    tag.send()
```