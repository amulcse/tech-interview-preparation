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
      $temp = $this->rootNode;

      while($temp){
        $previous = $temp;
        if($temp->node >= $data){
          $temp = $temp->left;
          if(!$temp){
            $previous->left = $newNode;
          }
        }
        else{
          $temp = $temp->right;
          if(!$temp){
            $previous->right = $newNode;
          }
        }
      }
    }
  }

  public function preOrderTraversal($node){
    if($node){
      echo ' '.$node->node;
      $this->preOrderTraversal($node->left);
      $this->preOrderTraversal($node->right);
    }
  }

  public function postOrderTraversal($node){
    if($node){
      $this->postOrderTraversal($node->left);
      $this->postOrderTraversal($node->right);
      echo ' '.$node->node;
    }
  }

  public function inorderTraversal($node){
    if($node){
      $this->inorderTraversal($node->left);
      echo ' '.$node->node;
      $this->inorderTraversal($node->right);
    }
  }

}


$binarySearchTree = new BinarySearchTree();
$binarySearchTree->insert(10);
$binarySearchTree->insert(14);
$binarySearchTree->insert(5);
$binarySearchTree->insert(1);
echo $binarySearchTree->inorderTraversal($binarySearchTree->rootNode)." </br> ";
echo $binarySearchTree->postOrderTraversal($binarySearchTree->rootNode)." </br> ";
echo $binarySearchTree->preOrderTraversal($binarySearchTree->rootNode)." </br> ";