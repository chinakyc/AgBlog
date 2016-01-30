angular.module("blogApp")
    .directive("myMoment", 
        function () {
            return {
                restrict: "E",
                scope: {
                    time: "=time",
                },
                // http://stackoverflow.com/questions/14300986/angularjs-directive-isolated-scope-and-attrs
                // https://docs.angularjs.org/guide/directive#creating-a-directive-that-manipulates-the-dom
                link: function (scope, elem, attr){
                    elem.text(moment(scope.time).lang("zh-cn")[attr.type]())
                },
            };
        }
    );
