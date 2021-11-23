<!-- 2.1 R􀂧mov􀂧 Dups! Write code to remove duplicates from an unsorted linked list. -->
<?php
include_once 'singled_linked_list.php';

$linkedList = new SingleLinkedList();

$linkedList->appendToHead('12');
$linkedList->appendToHead('25');
$linkedList->appendToHead('12');
$linkedList->appendToHead('12');
$linkedList->printList();
$linkedList->removeDuplicateWithoutBuffer();
$linkedList->printList();