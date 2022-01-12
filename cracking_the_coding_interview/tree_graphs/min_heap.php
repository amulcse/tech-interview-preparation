<?php

class MinHeap {

  function __construct(){
    $this->minHeap = [];
  }

  function insert($data){
    
    $this->minHeap[] = $data;
    $this->findRealPostion(sizeof($this->minHeap)-1);

  }

  function findRealPostion($i){

    $parentIndex = floor(($i-1)/2);
    while($parentIndex >= 0){
      if($this->minHeap[$parentIndex] > $this->minHeap[$i]){
        $temp =$this->minHeap[$parentIndex];
        $this->minHeap[$parentIndex] = $this->minHeap[$i];
        $this->minHeap[$i] = $temp;
        $i = $parentIndex;
        $parentIndex = floor(($i-1)/2);
      }
      else{
        break ;
      }
    }
    

  }

  function extractMin(){
    print_r($this->minHeap);
    return $this->minHeap[0];
  }

}

$minHeap = new MinHeap();
$minHeap->insert(4);
$minHeap->insert(7);
$minHeap->insert(3);
$minHeap->insert(1);
$minHeap->insert(-4);
$minHeap->insert(2);
$minHeap->insert(4);
echo $minHeap->extractMin();