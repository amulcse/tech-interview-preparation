import {
    Heap
} from 'heap-js';

// Min Heap by default
const minHeap = new Heap(Heap.maxComparator);

// Initialize the heap with an array
minHeap.init([5, 18, 1]);
// Push a new value
minHeap.push(2);

console.log(minHeap.pop()); //> 1
console.log(minHeap.pop()); //> 1
console.log(minHeap.pop()); //> 1
console.log(minHeap.pop()); //> 1