
$(function() {
    /**
     * issue的属性选择器value发生变化时
     * @param {*} id 
     * @param {*} type
     * @param {*} desc 
     * @param {*} extID
     */
    $.fn.fnDataPickerDidChangeValue = function (ID, type, desc, extID) {
        trackCallStack();
    
        setInnerHtmlWithID(ID + '-btn', desc);
        setHtmlValueWithID(ID, type);
    };
});
