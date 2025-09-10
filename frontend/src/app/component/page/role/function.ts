export function buildTreeForUI(data: any[]) {
  const map = new Map<string, any>()
  const tree: any[] = []

  // Initialize Map with transformed objects
  data.forEach(item => {
    map.set(item.id, { 
      title: item.name, 
      key: item.id, 
      children: [] 
    })
  })

  // Build hierarchical structure
  data.forEach(item => {
    const parentId = item.mainId
    if (parentId && map.has(parentId)) {
      // push the node object itself, not just id
      map.get(parentId)!.children.push(map.get(item.id))
    } else {
      tree.push(map.get(item.id))
    }
  })

  // Mark leaf nodes
  function markLeafNodes(nodes: any[]) {
    nodes.forEach(node => {
      if (node.children.length === 0) {
        node.isLeaf = true
      } else {
        markLeafNodes(node.children)
      }
    })
  }

  markLeafNodes(tree)

  return tree
}
