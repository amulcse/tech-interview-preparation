<?php
    function assert_failed($file, $line, $expr) {
        print "Assertion failed in $file on line $line: $expr\n";
    }

    assert_options (ASSERT_CALLBACK, 'assert_failed');
    assert_options (ASSERT_WARNING, 0);

    $foo = 10;
    $bar = 11;
    assert('$foo > $bar');
?>