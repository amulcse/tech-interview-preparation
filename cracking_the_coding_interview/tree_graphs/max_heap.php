<?php

class MaxHeap {

  public function __construct(){
    $this->maxHeap = [];
  }

  public function insert($data){
    $this->maxHeap[] = $data;
    $this->bubbleSwap(sizeof($this->maxHeap) - 1);
  }

  public function bubbleSwap($index){
    
    $parentIndex = floor(($index-1)/2);

    while($parentIndex >= 0){

      $parent = $this->maxHeap[$parentIndex];
      if($parent < $this->maxHeap[$index]){
        $temp = $this->maxHeap[$index];
        $this->maxHeap[$index] = $parent;
        $this->maxHeap[$parentIndex] = $temp;
        $index = $parentIndex;
        $parentIndex = floor(($parentIndex-1)/2);
      }
      else{
        break;
      }

    }

  }

  public function getMax(){
    return $this->maxHeap[0];
  }

  public function extractMax(){
    echo $this->maxHeap[0]."</br>";
    $lastIndex = sizeof($this->maxHeap)-1;
    $this->maxHeap[0] = $this->maxHeap[$lastIndex];
    unset($this->maxHeap[$lastIndex]);
    $this->bubbleSwapDown(0);
    print_r($this->maxHeap);
  }

  public function bubbleSwapDown($parentIndex){
    $totalElement = sizeof($this->maxHeap);
    while($parentIndex < $totalElement){

      $leftChildern = $parentIndex*2+1;
      $rightChildern = $parentIndex*2+2;

      $parent = $this->maxHeap[$parentIndex];

      if(isset($this->maxHeap[$leftChildern]) && $parent < $this->maxHeap[$leftChildern] &&  ( isset($this->maxHeap[$rightChildern]) && $this->maxHeap[$leftChildern] > $this->maxHeap[$rightChildern] )){
        $temp = $this->maxHeap[$leftChildern];
        $this->maxHeap[$leftChildern] = $parent;
        $this->maxHeap[$parentIndex] = $temp;
        $parentIndex = $leftChildern;
      }
      else if(isset($this->maxHeap[$rightChildern]) && $parent < $this->maxHeap[$rightChildern]){
        $temp = $this->maxHeap[$rightChildern];
        $this->maxHeap[$rightChildern] = $parent;
        $this->maxHeap[$parentIndex] = $temp;
        $parentIndex = $rightChildern;
      }
      else{
        break;
      }
    }
  }

}

$maxHeap = new MaxHeap();
$maxHeap->insert(3);
$maxHeap->insert(5);
$maxHeap->insert(89);
$maxHeap->insert(2);
$maxHeap->insert(45);
$maxHeap->insert(366);
$maxHeap->insert(-3);
$maxHeap->insert(500);
$maxHeap->extractMax();
$maxHeap->extractMax();
$maxHeap->extractMax();
$maxHeap->extractMax();