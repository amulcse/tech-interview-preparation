<?php

class BinaryArithmetic {

    private $number ;

    public function __construct($n)
    {   
        $this->number = bindec($n);
    }

    public function add($n){
        $this->number = $this->number+bindec($n);
    }

    public function getBinary(){
        return decbin($this->number);
    }

    public function addBinary($a, $b){
        return 
    }

}

$n = '1001';
$binaryArithmetic = new BinaryArithmetic($n);
$binaryArithmetic->add('01');
echo $binaryArithmetic->getBinary();

$bin1 = new BinaryArithmetic('101');
$bin2 = new BinaryArithmetic('011');
$bin1->addBinary($bin2);


// $binaryNumber = "";
// $k = floor($n/2);
// for($i = $k; $i >= 0 ; $i--){

//     $c = pow(2,$i);


//     if($n-$c >= 0){
//         $binaryNumber .= "1"; 
//         $n = $n - $c;
//     } else {
//         $binaryNumber .= "0";
//     }
// }

// echo "test ".$binaryNumber;