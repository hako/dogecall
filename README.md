```
                   
                                                                                           
                            .N.                                           lW. .W:      
                            :M.                                           oM. .M:      
                       ,kxxkKN  .xOO0l   :kOOOo  'kkxkk.  .kOxkd  ;kkxOx  dW  .M;      
                      ,N.   kK .Wo  .Wc 0O.  cW.'N,'lkk. .Nc     oX.  dK  dN  .M;      
                      cX    xK ;M,   Nc.Mc  .KK lWxc.  . lX      Ok   l0  kX  'M,      
                       dKxxOKN  oKxxKo  ;0OOx0k  d0xdxkx  k0dokx 'KkldOW; k0  ,M.      
                          .       ..        .Nl     .       ..     ..                  
                                        dxxx0d                                                                                                                 
```
###dogecall

`dogecall` is a port of my python program `dogecall.py` which was used for [this silly video](https://www.youtube.com/watch?v=9S3BX62vToo) I made a year ago. It was made as a joke remaking [this post I stumbled upon](http://hakob.yt/doge). It was fun doing it.

Installation
-----
dogecall comes as two programs:

`dogecall` - run on PC.

`dogecall-server` - run on a server.

To install (If you have Go installed)

`go get github.com/hako/dogecall`

`go get github.com/hako/dogecall/cmd/dogecall-server`

Usage
-----

dogecall only requires you to have a Twilio account. If you would like to use dogecall, please [create an account](https://www.twilio.com/try-twilio) and come back.

These can be set in a file called `dogecall.rc` when you install the `dogecall` command.


`dogecall` Usage:
----------
```
dogecall - ENCOUNTER A DOGE THROUGH A PHONE CALL.

Usage:
 dogecall [-n <number>]
 dogecall [-h | --help] [-v | --version]

Options:
  -n [PHONE NUMBER]   such fone numer (with area codez) eg. +44 = 44
  -h, --help          show this help message and exit.
  -v, --version       the versions lol.
```

When you run `dogecall` for the first time, you have to configure `dogecall.rc`.


AccountSid = Your Twilio Account SID. (Found in your account dashboard)

URL = `http://dc.hakobaito.co.uk` (`dogecall-server` Default URL.)

TwNumber = Your Twilio Phone number. (Found when you get a Twilio number.)

TwAuthtoken = Your Twilio auth token. (Found in your account dashboard.)


`dogecall-server` Usage:
----------

```
dogecall-server - SERVER FOR DOGECALL.

Usage:
 dogecall [-s]
 dogecall [-h | --help] [-v | --version]

Options:
  -s                  wow lanch server 2 serve u.
  -h, --help          show this help message and exit.
  -v, --version       the versions lol.
```

If you intend to deploy `dogecall-server` don't forget to set the **config variables** on your hosting platform.

dogecall-server default environment variables:

`PORT`   - Whatever you like (ie, 8080)

`GO_ENV` - `production` (set only on a production server)

License
-----
MIT
