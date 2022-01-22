<?php

Class DistinctUnionIterator implements Iterator {

  private $head = 0;
  private $uniqueResults = [];
  private $curretIterator ;

  public function __construct(Iterator $sourceA, Iterator $sourceB) {
      $this->position = 0;
      $this->curretIterator = $sourceA;
      $this->sourceA = $sourceA;
      $this->sourceB = $sourceB;
  }

  public function rewind() {
      $this->head = $this->curretIterator->rewind();
  }

  public function current() {
    while(in_array($this->curretIterator->current(),$this->uniqueResults)){
      $this->next();
    }
    $current = $this->curretIterator->current();
    $this->uniqueResults[] = $current;
    return $current;
  }

  public function key() {
    return $this->head;
  }

  public function next() {
    $this->head = $this->curretIterator->next();
  }

  public function valid() {
    if($this->sourceA->valid()){
      $this->curretIterator = $this->sourceA;
      return true;
    }
    else if($this->sourceB->valid()){
      $this->curretIterator = $this->sourceB;
      return true;
    }
    else{
      return false;
    }
  }
}

$a = [1,2];
$aIterator = new ArrayIterator($a);

$b = [0,3,6,4,7,5];
$bIterator = new ArrayIterator($b);

$it = new DistinctUnionIterator($aIterator, $bIterator) ;

foreach($it as $key => $value) {
  echo $value."\n";
}
?>