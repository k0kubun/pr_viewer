jQuery(function($) {
  $(".navbar-search").submit(function() {
    var username = $('.search-query').val();
    location.href = "/users/" + username;
    return false;
  });
});
