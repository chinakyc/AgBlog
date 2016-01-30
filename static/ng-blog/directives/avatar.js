angular.module("blogApp")
    .directive("myAvatar", 
        function () {
            return {
                restrict: "A",
                scope: {
                    email: "=email",
                },
                // http://stackoverflow.com/questions/14300986/angularjs-directive-isolated-scope-and-attrs
                // https://docs.angularjs.org/guide/directive#creating-a-directive-that-manipulates-the-dom
                link: function (scope, elem, attr){
                          elem.attr("src", "http://cn.gravatar.com/avatar/" + md5(scope.email) + "?d=identicon&s=" + attr.size)
                },
            };
        }
    );
