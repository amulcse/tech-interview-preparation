<?php

class DoubleIterator implements Iterator {

    private $items = [];
    private $pointer = 0;

    public function __construct($items){
        $this->items = array_values($items);
    }

    public function current():mixed {
        return "hello".$this->items[$this->pointer];
    }

    public function key():mixed {
        return $this->pointer;
    }

    public function next():void {
        $this->pointer++;
    }

    public function rewind():void  {
        $this->pointer = 0;
    }

    public function valid():bool {
         return $this->pointer < count($this->items);
    }

}

$ii = new DoubleIterator([1,2,4,5,6,6]);

foreach($ii as $key => $value){
    print_r($key." ".$value." \n");
}