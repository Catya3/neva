import { Edge, Node } from "reactflow";
import ELK, { ElkNode } from "elkjs";

const nodeWidth = 350;
const nodeHeight = 100;

const elk = new ELK();

export default async function getLayoutedNodes(
  nodes: Node[],
  edges: Edge[]
): Promise<Node[]> {
  const graph: ElkNode = {
    id: "root",
    layoutOptions: {
      "elk.algorithm": "mrtree",
      "elk.direction": "DOWN",
    },
    children: ["type", "const", "interface", "component"].map(
      (groupNodeId) => ({
        id: groupNodeId,
        width: nodeWidth,
        height: nodeHeight,
        layoutOptions: { "elk.direction": "DOWN" },
        children: nodes
          .filter((node) => node.type === groupNodeId)
          .map((node) => ({
            ...node,
            width: nodeWidth,
            height: nodeHeight,
            layoutOptions: {
              "elk.direction": "DOWN",
            },
          })),
      })
    ),
    edges: edges.map((edge) => ({
      id: edge.id,
      sources: [edge.source],
      targets: [edge.target],
    })),
  };

  const layoutedGraph = await elk.layout(graph);

  const layoutedNodes: Node[] = [];
  for (const groupNode of layoutedGraph.children!) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const data = (groupNode as any).data;
    layoutedNodes.push({
      ...groupNode,
      id: groupNode.id,
      position: { x: groupNode.x!, y: groupNode.y! },
      data: { ...data, label: groupNode.id },
      style: { width: groupNode.width, height: groupNode.height },
    });
    for (const childNode of groupNode.children!) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const data = (childNode as any).data;
      layoutedNodes.push({
        ...childNode,
        id: childNode.id,
        position: { x: childNode.x!, y: childNode.y! },
        data: { ...data, label: childNode.id },
        style: { width: childNode.width, height: childNode.height },
        parentNode: groupNode.id,
      });
    }
  }

  return layoutedNodes;
}
