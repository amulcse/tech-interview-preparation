<?php

class Node {

  public $data;
  public $link;

  function __constuct($data = '', $link = null){
    $this->data = $data;
    $this->link = $link;
  }

}

class SingleLinkedList {

  public $head ;
  public $tail ;

  function __construct(){
    $this->head  = null;
    $this->tail = null;
  }

  /* Append data to tail */
  public function appendToTail($data){

    $node = new Node($data);
    $node->data = $data;

    if($this->tail){
      $this->tail->link = $node;
    }

    /* Update current node */
    $this->tail = $node;

    /* Initialise Head */
    if($this->head == null){
      $this->head = $node;
    }

  }

  /* Apend data to Head */
  public function appendToHead($data){
    $node = new Node($data);
    $node->data = $data;

    if($this->head){
      $node->link = $this->head;
    }

    $this->head = $node;
  }

  public function printList(){
    $traverseNode = new Node();
    $traverseNode = $this->head;
    if($traverseNode){
      while($traverseNode){
        echo $traverseNode->data."\n";
        $traverseNode = $traverseNode->link;
      }
    }
  }

  public function deleteNode($data){
    $currentNode = new Node();
    $previousNode = null;
    $currentNode = $this->head;
    $foundKey = false;
    while($currentNode && !$foundKey){
      if($currentNode->data == $data){
        /* first element */
        if(!$previousNode){
          $this->head = $this->head->link;
        }
        else{
          $previousNode->link = $currentNode->link;
        }
        $foundKey = true;
      }
      $previousNode = $currentNode;
      $currentNode = $currentNode->link;
    }
  }

  public function getNodeCounts(){
    if(!$this->head){
      return 0;
    }
    $counter = 0;

    $current = new Node();
    $current = $this->head;

    while($current){
      $counter++;
      $current = $current->link;
    }

    return $counter;
  }

  public function search($key,$current){
    if(!$current){
      return 'No';
    }
    else if($current->data == $key){
      return 'Yes';
    }
    else{
      return $this->search($key,$current->link);
    }
  }

  public function removeDuplidate(){
    $current = new Node();
    $current = $this->head;
    $previous = null;
    $visitedKeys = [];
    while($current){
      if(!isset($visitedKeys[$current->data])){
        $visitedKeys[$current->data] = 1;
      }
      else{
        $previous->link = $current->link;
      }
      $previous = $current;
      $current = $current->link;
    }
  }

  public function removeDuplicateWithoutBuffer(){
    $traverseNode = new Node();
    $traverseNode = $this->head;
    if($traverseNode){
      while($traverseNode){
            $current = new Node();
            $current = $traverseNode->link;
            $previous = $traverseNode;
            while($current){
              if($traverseNode->data === $current->data){
                $previous->link = $current->link;
              }
              $previous = $current;
              $current = $current->link;
            }
            $traverseNode = $traverseNode->link;
      }
    }
  }

}

$singleLinkedList = new SingleLinkedList();

$singleLinkedList->appendToHead('India');
$singleLinkedList->appendToHead('Canada');
$singleLinkedList->appendToHead('USA');
$singleLinkedList->appendToHead('Japan');
// $singleLinkedList->printList();

// $singleLinkedList->deleteNode('Canada');
// $singleLinkedList->printList();

// echo "Total node in Linked list is : ".$singleLinkedList->getNodeCounts();

// echo " Is India in linked list : ".$singleLinkedList->search('India',$singleLinkedList->head);

