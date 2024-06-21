#!/usr/bin/env zsh

CURDIR=$(cd $(dirname $0); pwd)
BinaryName=nobot
echo "$CURDIR/bin/${BinaryName}"

# plz init your database when you first run this script
# $CURDIR/bin/${BinaryName} twt init

#fetch users data from twitter

while true
  do
    $CURDIR/bin/${BinaryName} twt fetch-save -c 5
    #fetch users data from nostr
    $CURDIR/bin/${BinaryName} twt nf -r wss://wss://relay.snort.social \
    -r wss://nos.lol \
    -r wss://offchain.pub \
    -r wss://nostr-pub.wellorder.net \
    -r wss://relay.damus.io \
    -r wss://freerelay.xyz \
    -r wss://relay.current.fyi \
    --pubkey npub1z0nsnf2gwm7ydfj8gxq5v5fhcljarx0waf6rz6zhy4gtrmaevvtqc3qsun \
    --pubkey npub1whqanghnxln3cqc78lrvtlmta4sf9vq04awq3spa9gryly9ae6ks06xrzd \
    --pubkey npub1pw5fs3m5xkw693v6fkmace9x8mdhmmvmxczeq77ayavl6yyfxmfq096yuz \
    --pubkey npub1mzntze6y4vnvjhh3rxhuctfa3eqq5cxse9ql9g8f2lkztsxd2zxqlgm5ru \
    --pubkey npub1l494xy07jwv4gwp0yjjly9myt0yhm4d5n6d5uh4nlt46qsrd223q9kd5sy \
    --pubkey npub1rr0eajwdr7jwfad9rz3dwtdcu38yenxfke40ngrsdx32384fxnyqa293qe \
    --pubkey npub19vc0ly5dkhpu5jfvj72vkkngfaa75fwcppkn9a6nj89tylj2qrfqrl3vgu \
    --pubkey npub1eatx8rjaa0nllttrwqe2csawg7pyvjdayyd89vch8jp6552z6hxqhphnvn

    #output daily report
    $CURDIR/bin/${BinaryName} twt dr
  done