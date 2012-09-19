Blunderbluss
============

Simple example of golang code with a blunderbluss pattern.

Telnet shot, http clients get lead.

Build
-----

    ./build.sh

Test it
-------

    ./blunderbluss

Lauch some curls :

    curl http://localhost:8000/blunder

Send some events :

    telnet localhost 5000

Every lines in the telnet session is broadcasted to each http connections.

Todo
----

Some javascript nice examples.

Licence
-------

Three terms BSD.
