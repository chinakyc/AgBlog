angular.module("blogApp")
    .directive("messageBox", ["Session", "MESSAGE_EVEN", "MESSAGE_TYPE",
        function (Session, MESSAGE_EVEN, MESSAGE_TYPE) {
            return {
                restrict: "A",
                templateUrl: "templates/message-box.html",
                link: function (scope) {
                    scope.messages = [];
                    scope.messageType = MESSAGE_TYPE;
                    scope.visiable = false;

                    function showBox () {
                        scope.visiable = true;
                        // get and clean messages in Session
                        scope.messages = Session.getMessages();
                    };

                    //listen event of flashed
                    scope.$on(MESSAGE_EVEN.flashed, showBox);
                }
            };
        }
    ]);
