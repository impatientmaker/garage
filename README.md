# garage
REST API to call USB relay and open my garage door (https://github.com/darrylb123/usbrelay)

This is basically a backup for my own use.  But feel free to use it.

Installing drivers:

This info is derived from this page: https://github.com/darrylb123/usbrelay

That package depends on HID API, which I installed via apt.
(this is on an ARM platform, namely https://getchip.com/pages/chip)

On chip, which is debian-based, you can install the HID API packages thusly:
# apt-get install libhidapi-dev libhidapi-libusb0 build-essential

#   #perform the udev rules step recommended (although it doesn't seem to work)
so screw it, using a hammer:
in /etc/rc.local:
chmod 666 /dev/hidraw0 



Don't forget to install Go (arm version).
