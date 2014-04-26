jQuery(function($) {
  $(".navbar-search").submit(function() {
    var username = $('.search-query').val();
    location.href = "/users/" + username;
    return false;
  });

  $(".tab-change").click(function(event) {
    event.preventDefault();
    $(".tab-content").removeClass("active");
    $(".nav-tabs li").removeClass("active");

    tabLink = $(event.target).closest("a");
    tab = tabLink.closest("li");
    tab.addClass("active");

    content = $(tabLink.attr("id"));
    content.addClass("active");
  });
});
