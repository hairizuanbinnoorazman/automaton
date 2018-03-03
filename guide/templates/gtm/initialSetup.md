# GTM Implementation Guide

## Initial Setup

Add the following script script before the GTM Code snippet below.

```html
<script>
    var datalayer = datalayer || [];
</script>
```

Copy the following JavaScript and paste it as close to the opening `<head>` tag as possible on every page of your website.
It should as close to the top of the html page as possible so that the snippet can get loaded as soon as possible before events that are relevent to your business' analytics get triggered.

```html
<!-- Google Tag Manager -->
<script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
})(window,document,'script','dataLayer','{{.GtmContainerID}}');</script>
<!-- End Google Tag Manager -->
<!-- Google Tag Manager (noscript) -->
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id={{.GtmContainerID}}"
height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
<!-- End Google Tag Manager (noscript) -->
```