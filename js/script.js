var script_second = 0;
var main_part = "";
var locate_menu = null;

function script_start() {
    pool_step.push(script_step);
    if (!lik_trust) lik_set_trust("");
    if (lik_trust) lik_set_marshal(1000, "/front/marshal");
    topon_init();
}

function script_step() {
    script_showtime();
    script_redraw();
}

function script_showtime() {
    if (script_second!=tick_second) {
        script_second = tick_second;
        var elm = jQuery('#srvtime');
        var tt = tick_server - tick_shift_minute * 60;
        var ts = tt % 60;
        tt = (tt - ts) / 60;
        var tm = tt % 60;
        tt = (tt - tm) / 60;
        var th = tt % 24;
        tt = (tt - th) / 24;
        var text = (th >= 10) ? "" + th : "0" + th;
        text += (ts & 1) ? ":" : " ";
        text += (tm >= 10) ? tm : "0" + tm;
        text += (ts>=10) ? ":"+ts : ":0"+ts;
        elm.text(text);
    }
}

function script_redraw() {
    let rdr = jQuery('[redraw]');
    if (rdr.size() > 0) {
        rdr.each(function (idx, item) {
            let elm = jQuery(item);
            let redraw = elm.attr('redraw');
            elm.removeAttr('redraw');
            setTimeout(function () {
                if (redraw in window) {
                    window[redraw](elm);
                }
            }, 50);
        });
    }
}

function combine_front(part) {
    let path = "/front";
    if (/^\//.exec(part)) {
        path += part;
    } else if (part) {
        path += "/" + part;
    }
    return path;
}

function front_get(cmd) {
    get_data_part(combine_front(cmd));
}

function front_post(cmd, data) {
    post_data_part(combine_front(cmd), data);
}

function front_proc(cmd, proc, parm) {
    get_data_proc(combine_front(cmd), proc, parm);
}

function front_post_proc(cmd, data, proc, parm) {
    post_data_proc(combine_front(cmd), data, proc, parm);
}

function click_set_part(part) {
    get_data_part("/front/part" + part);
}

function grid_id_cmd(id, cmd) {
    click_button(id);
    get_data_part("/front/cmd/" + cmd);
}

function run_syscmd(cmd) {
    get_data_part("/front/system/" + cmd);
}

function grid_set_sort(sort) {
    get_data_part("/front/sort/" + sort);
}

function grid_rule_list(cmd) {
    get_data_part("/front/page/" + cmd);
}

function grid_set_seek(nelm,idb) {
    jQuery(".grid_inf").removeClass("grid_sel");
    jQuery(".grid_row"+nelm).addClass("grid_sel");
    jQuery("#grid_num").text(nelm+1);
    get_data_part("/front/seek/" + idb);
}

function grid_click_seek(id) {
    get_data_part("/front/deck/" + id);
}

function grid_mouse_over(nelm, ncol) {
    //alert("Over: " + nelm + "," + ncol);
    jQuery(".grid_row"+nelm).filter(".grid_col"+ncol).addClass("grid_loc");
    //jQuery(".grid_col"+ncol).addClass("grid_loc");
}

function grid_mouse_out(nelm, ncol) {
    //alert("Out: " + nelm + "," + ncol);
    jQuery(".grid_inf").removeClass("grid_loc");
}

function grid_set_sector(field, value, oldValue) {
    do_cmd("/sector/" + value);
}

function grid_set_locate(field, value, oldValue) {
    do_cmd("/locate/" + value);
}

function deck_id_cmd(id, cmd) {
    click_button(id);
    get_data_part("/front/cmd/" + cmd);
}

function deck_id_path(id, cmd, path) {
    click_button(id);
    var data = edit_collect();
    post_data_part("/front/cmd/" + cmd + "/" + path, data);
}

function pop_id_choose(id, part) {
    click_button(id);
    menu_close_all();
    var over = jQuery("#"+id).closest(".popover");
    if (over.size()>0) {
        locate_menu = mouse_locate;
        var path = over.attr('rulepath');
        get_data_proc("/front/choose/"+part+"/"+path, choose_menu, over);
    }
}

var pop_menu;

function choose_menu(over, lika) {
    var code = ('code' in lika) ? lika.code : '';
    if (!code) {
        menu_close_all();
    } else {
        var level = over.parents(".popin").size();
        menu_close_level(level);
        menu_open(over, "<div class=popin>" + code + "</div>");
        if (locate_menu) {
            var pos = {left: locate_menu.X, top: locate_menu.Y};
            delete (locate_menu);
            pop_menu = over.find(".popin");
            pop_menu.offset(pos);
            pop_menu.css('z-index', 9999);
            setTimeout(align_menu, 500);
        }
        //menu_ready();
    }
}

function align_menu() {
    var offset = pop_menu.offset();
    if (offset.left + pop_menu.width() > document.body.clientWidth - 4) {
        offset.left = document.body.clientWidth - 4 - pop_menu.width();
    }
    if (offset.top + pop_menu.height() > document.body.clientHeight - 4) {
        offset.top = document.body.clientHeight - 4 - pop_menu.height();
    }
    pop_menu.offset(offset);
}

function deck_choose_elm(sys, text, path) {
    menu_close_all();
    jQuery("#edit_"+path).val(sys);
    jQuery("#show_"+path).val(text);
}

function item_id_cmd(cmd) {
    click_button(cmd);
    get_data_part("/front/cmd/" + cmd);
}

function item_id_save(cmd) {
    click_button(cmd);
    var data = edit_collect();
    post_data_part("/front/cmd/" + cmd, data);
}

function click_button(id) {
    jQuery("#" + id).addClass("picon");
    setTimeout(function(){ jQuery(".picon").removeClass("picon"); }, 250);
}

function edit_collect() {
    var data = "";
    jQuery("[id^=edit_]").each(function(idx,item) {
        var elm = jQuery(item);
        var key = elm.attr('id');
        var val = elm.val();
        if (elm.attr('type')=='checkbox') {
            val = (elm.prop('checked')) ? "1" : "0"
        }
        if (data) data += "&";
        data += key + "=";
        data += string_to_XS(val);
    });
    return data;
}

function deck_redraw(height) {
    get_data_part("/front/decking/" + height);
}

function tsan_key_press(e) {
    //alert('key');
    var keycode = (e.keyCode ? e.keyCode : e.which);
    if (keycode == '13') {
        //alert('You pressed enter! - keypress');
    }
}

function set_choose(set) {
    get_data_part("/front/cmd/area_setchoose/" + set);
}

function show_progress(item) {
    var checked = (jQuery(item).prop("checked")) ? 1 : 0;
    get_data_part("/front/cmd/area_hideprogress/" + checked);
}

function tune_control(item) {
    var checked = (jQuery(item).prop("checked")) ? 1 : 0;
    get_data_part("/front/cmd/area_attune/" + checked);
}

function album_click(grid, obj) {
    var id = obj.id;
    var cmd = obj.column.index;
    if (cmd=="map" || cmd=="scheme") {
        var val = (obj.value) ? 0 : 1;
        do_cmd("album_"+cmd + "/" + id + "_" + val);
    }
}

function row_delete(grid, params) {
    do_cmd("album_delete/" + params.id + "_0");
}

function dragrows(grid, params) {
    var it = params[0];
    var id = it.id;
    var data = grid.getData();
    for (var nr = 0; nr < data.length; nr++) {
        var row = data[nr];
        if (id == row.id) {
            do_cmd("album_order/" + id + "_" + nr);
            break;
        }
    }
}

function dump(params) {
    var text = "";
    for (var key in params) {
        text += key + ": " + params[key] + "\n";
    }
    alert(text);
}

function open_bind_offer(idoff) {
    fancy_window_destroy();
    lik_window_part("/show/offer"+idoff+"?_tp=1");
    window.location.reload();
}

function open_bind_offer(idoff) {
    fancy_window_destroy();
    lik_window_part("/show/offer"+idoff+"?_tp=1");
    window.location.reload();
}

function open_create_offer(dir,fwr) {
    if (fwr > 0) {
        var form = FancyWidowForm;
        var data = "likid=" + form.likId;
        var values = form.get();
        for (var key in values) {
            data += "&edit_" + key + "=" + string_to_XS(values[key]);
        }
        fancy_window_destroy();
        post_data_part("/front/cmd/bell_addoffer/"+dir, data);
    } else {
        fancy_window_destroy();
        get_data_part("/front/cmd/bell_addoffer/"+dir);
    }
}

var setFixValue = "";

function set_choose(field, value, oldValue) {
    get_data_part("/front/cmd/_setchoose/" + value);
}

function set_duplicate() {
    get_data_part("/front/cmd/_setclone");
}

function set_delete() {
    get_data_part("/front/cmd/_setdelete");
}

function set_fixname(grid, value) {
    setFixValue = value;
}

function set_rename() {
    get_data_part("/front/cmd/area_setrename/" + string_to_XS(setFixValue));
}

function bell_set_change(obj) {
    let form = fancy_seekform(obj);
    form.get("phone1");
}

function bell_accept(ok) {
    get_data_part("/front/cmd/all_accept/" + ok);
    setTimeout(form_reshow, 500);
}

function form_reshow() {
    fancy_window_destroy();
    fancy_bind_form("all", "_show");
}

function fullScreen() {
    var element = document.documentElement;
    if(element.requestFullscreen) {
        element.requestFullscreen();
    } else if(element.webkitrequestFullscreen) {
        element.webkitRequestFullscreen();
    } else if(element.mozRequestFullscreen) {
        element.mozRequestFullScreen();
    }
}

function topon_init() {
    $(function () {
        $('#topon a').stop().animate({'marginLeft':'-45px'},250);
        $('#topon > li').hover(
            function () {
                $('a', $(this)).stop().animate({'marginLeft': '5px'}, 200);
            },
            function () {
                $('a', $(this)).stop().animate({'marginLeft': '-45px'}, 200);
            }
        );
    });
}

function topon_click() {
    front_get("/admin");
}

function tune_canal(code, value) {
    let answer = prompt("Новое значение:", value)
    if (answer !== false && answer != value) {
        front_get("/tune/" + code + "/" + string_to_XS(answer));
    }
}

function tune_canal_code() {
    let elm = jQuery("#canalcode");
    if (elm) {
        let code = elm.val();
        front_get("/tune/canalcode/" + string_to_XS(code));
    }
}

function tune_append() {
    let answer = prompt("Укажите код нового канала:", "")
    if (answer !== false && answer != "") {
        front_get("/tune/append/" + string_to_XS(answer));
    }
}

function tune_delete(idc) {
    if (confirm("Подтвердите удаление канала " + idc)) {
        front_get("/tune/delete");
    }
}

function tune_edit_start() {
    front_get("/tune/editstart");
}

function tune_edit_write() {
    front_get("/tune/editwrite");
}

function tune_edit_cancel() {
    front_get("/tune/editcancel");
}

