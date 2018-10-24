$(function () {
    /**
     * issue的属性选择器value发生变化时
     * @param {*} id 
     * @param {*} type
     * @param {*} desc 
     * @param {*} extID
     */
    $.fn.fnDataPickerDidChangeValue = function (ID, type, desc, extID, callback) {
        
        setInnerHtmlWithID(ID + '-btn', desc);
        setHtmlValueWithID(ID, type);

        if (callback.length > 0 && callback != "undefined") {
            var fn = eval("$()." + callback);
            new fn(ID, type, extID);
        }
    };
});