<?php

class Node {
  public $node = null;
  function __construct($data){
      $this->node = $data;
      $this->left = null;
      $this->right = null;
  }
}

class BinarySearchTree {

  public $bTree = null;
  public $rootNode = null;
  function __construct()
  {
    $this->bTree = null;
    $this->rootNode = null;
  }

  public function insert($data){
    $newNode = new Node($data);

    if(!$this->rootNode){
      $this->rootNode = $newNode;
    }
    else{
      $temp = $previous = $this->rootNode;

      while($temp->node){

        $previous = $temp;
        if($temp->node >= $data){
          $temp = $temp->node->left;
          if(!$temp){
            $previous->left = $newNode;
          }
        }
        else{
          $temp = $temp->node->right;
          if(!$temp){
            $previous->right = $newNode;
          }
        }

      }
    }
  }

  public function inorderTraversal($node){
    if($node){
      $this->inorderTraversal($node->left);
      echo $this->inorderTraversal($node->node);
      $this->inorderTraversal($node->right);
    }
  }

}


$binarySearchTree = new BinarySearchTree();
$binarySearchTree->insert(10);
$binarySearchTree->insert(14);
$binarySearchTree->insert(5);
$binarySearchTree->insert(1);
$binarySearchTree->inorderTraversal($binarySearchTree->rootNode);