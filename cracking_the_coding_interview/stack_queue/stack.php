<?php

class Stack {

  public $stack ;
  public $min = 0;

  function __construct(){
    $this->stack = [];
  }

  public function pop(){
    $poped = array_pop($this->stack);
    echo $poped." is popped from stack";
    if($this->min  $data){
      $this->min = $data;
    }
    return $poped;
  }

  public function push($data){
    echo $data." pushed to stack";
    array_push($this->stack,$data);
    if($this->min > $data){
      $this->min = $data;
    }
  }

  public function isEmpty(){
    return (sizeof($this->stack) > 0)? false : true;
  }

  public function peek(){
    return ($this->isEmpty())? -1 : $this->stack[sizeof($this->stack)-1];
  }

  public function min(){
    return $this->min;
  }

}

$stack = new Stack();
$stack->push('1');
$stack->push('2');
$stack->push('3');
$stack->push('4');
echo $stack->peek();
echo $stack->pop();
echo $stack->peek();
