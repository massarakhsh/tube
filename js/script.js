var script_second = 0;

function script_start() {
    pool_step.push(script_step);
    if (!lik_trust) lik_set_trust("");
    if (lik_trust) lik_set_marshal(3000, "/front/marshal");
    tube_init();
}

function script_step() {
    script_showtime();
    script_redraw();
}

function script_showtime() {
    if (script_second!=tick_second) {
        script_second = tick_second;
        var elm = jQuery('#srvtime');
        if (elm.size()>0) {
            var text = "";
            var ok = true;
            if (tick_total - tick_answer < 5*1000) {
                var tt = tick_server - tick_shift_minute * 60;
                text = build_showtime(tt, true);
                ok = true;
            } else if (tick_total - tick_answer < 5*60*1000) {
                var tt = Math.floor((tick_total - tick_answer) / 1000);
                text = "Нет связи " + build_showtime(tt, false);
                ok = false;
            } else {
                lik_stop();
                text = "<b>СИСТЕМА ОСТАНОВЛЕНА</b>";
                ok = false;
            }
            if (ok && elm.hasClass("srvoff")) {
                elm.removeClass("srvoff");
            } else if (!ok && !elm.hasClass("srvoff")) {
                elm.addClass("srvoff");
            }
            elm.html(text);
        }
    }
}

function build_showtime(tt, ok) {
    var ts = tt % 60;
    tt = (tt - ts) / 60;
    var tm = tt % 60;
    tt = (tt - tm) / 60;
    var th = tt % 24;
    tt = (tt - th) / 24;
    var text = "";
    if (ok) {
        text += (th >= 10) ? "" + th : "0" + th;
        text += (ts & 1) ? ":" : " ";
        text += (tm >= 10) ? tm : "0" + tm;
    } else {
        if (th > 0) text += ""+th+":";
        text += tm;
        text += (ts >= 10) ? ":" + ts : ":0" + ts;
    }
    return text;
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

function edit_write(cmd) {
    var data = edit_collect();
    front_post(cmd, data);
}

function edit_collect() {
    let data = "";
    jQuery('[id^=up_]').each(function(idx,item) {
        var elm = jQuery(item);
        if (data) data += "&";
        data += elm.attr('id') + "=" + string_to_XS(elm.val());
    });
    return data;
}

function tube_init() {
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

function tube_command(fun) {
    let path = fun;
    if (fun == 'write') {
        path = '/admin/write';
    } else if (fun == 'cancel') {
        path = '/admin/cancel';
    } else if (fun == 'edit') {
        path = '/admin/edit';
    } else if (fun == 'copy') {
        path = '/admin/copy';
    } else if (fun == 'delete') {
        if (!confirm("Подтвердите удаление канала")) {
            return;
        }
        path = '/admin/delete';
    } else if (fun == 'create') {
        path = '/admin/create';
    } else {
    }
    if (path != /^\//) {
        path = '/' + path;
    }
    front_get(path);
}

function tube_format(item) {
    let answer = jQuery(item).val();
    front_get('/admin/format/' + string_to_XS(answer));
}

function tube_code(value) {
    let answer = prompt("Новый код:", value)
    if (answer !== false && answer != value) {
        front_get("/admin/code/" + string_to_XS(answer));
    }
}

function tube_name(value) {
    let answer = prompt("Новое наименование:", value)
    if (answer !== false && answer != value) {
        front_get("/admin/name/" + string_to_XS(answer));
    }
}

function tube_source(ns) {
    front_get("/admin/source/" + ns);
}

function media_path(path) {
    front_get("/media/path/" + string_to_XS(path));
}

function media_file(path) {
    front_get("/media/file/" + string_to_XS(path));
}

function media_load() {
    front_get("/media/load");
}

function media_delete() {
    front_get("/media/delete");
}

function media_store() {
    front_get("/media/store");
}

function media_cancel() {
    front_get("/media/cancel");
}

