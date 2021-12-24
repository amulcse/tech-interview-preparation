function solution($numbers) {
    $winnersName = ['Alice','Bob'];
    $winner = 1;
    $maxIterations = sizeof($numbers);
    for($i=0;$i<$maxIterations-1;$i++){
        for($j=0; $j<sizeof($numbers)-1;$j++){
            if($numbers[$j] == $numbers[$j+1]){            
                $winner = ($winner+1)%2;
                array_splice($numbers,$j,2);
                break;
            }        
        }
        if(sizeof($numbers) <= 1){
            break;
        }
    }
    return $winnersName[$winner];
}
