# HEAP
---
## Defenetion 
• The value of each node must be greater than each of its descendant nodes.
This rule is known as the heap condition.
• The tree must be complete. (I’ll explain the meaning of this shortly.)
---
## Heep Insertion
1. We create a node containing the new value and insert it at the next
available rightmost spot in the bottom level. Thus, this value becomes
the heap’s last node.
2. Next, we compare this new node with its parent node.
3. If the new node is greater than its parent node, we swap the new node
with the parent node.
4. We repeat Step 3, effectively moving the new node up through the heap,
until it has a parent whose value is greater than it.
---
## Heap Deletion
1. Move the last node into where the root node was, effectively removing the
original root node.
2. Trickle the root node down into its proper place. I’ll explain how trickling
down works shortly. 
### rule to swap after deletion:
1. We check both children of the trickle node and see which one is larger.
2. If the trickle node is smaller than the larger of the two child nodes, we
swap the trickle node with that larger child.
3. We repeat Steps 1 and 2 until the trickle node has no children who are
greater than it.

