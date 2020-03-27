$browser.get('http://example.com')
.then(function() {
  return $browser.waitForPendingRequests();
});