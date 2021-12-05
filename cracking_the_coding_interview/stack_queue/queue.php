<?php

class Node {
  
  function __consturct($data){
    $this->data = $data;
    $this->next = null;
  }

}
Class Queue {

  public $queue;
  public $rear = null;
  public $front = null;

  function __construct(){
    $this->queue = [];
  }

  public function enqueue($data){

    $node = new Node($data);
    $node->data = $data;
    if($this->rear){
      $this->rear->next = $node;
    }
    $this->rear = $node;

    if($this->front == null){
      $this->front = $node;
    }

    echo $data." is enqued. ";

  }

  public function dequeue(){
    if($this->front){
      $front = $this->front;
      $this->front = $this->front->next;

      if($this->front == null){
        $this->rear = null;
      }

      echo $front->data." is deenqued. ";

      return $front;
    }
    else{
      echo "Queue is empty";
      return ;
    }
  }

  public function rear(){
    return $this->rear;
  }

  public function front(){
    return $this->front;
  }

}

$queue = new Queue();
$queue->enqueue('20');
$queue->enqueue('30');
$queue->enqueue('40');
$queue->dequeue();
$queue->enqueue('50');
$queue->dequeue();