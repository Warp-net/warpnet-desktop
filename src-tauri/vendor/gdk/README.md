# Rust GDK 3 bindings

**This project is UNMAINTAINED. Please take a look at [gtk4-rs](https://github.com/gtk-rs/gtk4-rs) instead!**

Project site is [here](https://gtk-rs.org/).

__Rust__ bindings and wrappers for [GDK 3](https://developer.gnome.org/gdk3/),
part of [gtk3-rs](https://github.com/gtk-rs/gtk3-rs).

GDK __3.22__ is the lowest supported version for the underlying library.

## Minimum supported Rust version

Currently, the minimum supported Rust version is `1.70.0`.

## Documentation

 * [Rust API - Stable](https://gtk-rs.org/gtk3-rs/stable/latest/docs/gdk/)
 * [Rust API - Development](https://gtk-rs.org/gtk3-rs/git/docs/gdk)
 * [C API](https://developer.gnome.org/gdk3/stable/)
 * [GTK Installation instructions](https://www.gtk.org/docs/installations/)

## Using

We recommend using [crates from crates.io](https://crates.io/keywords/gtk-rs),
as [demonstrated here](https://gtk-rs.org/#using).

If you want to track the bleeding edge, use the git dependency instead:

```toml
[dependencies]
gdk = { git = "https://github.com/gtk-rs/gtk3-rs.git", package = "gdk" }
```

Avoid mixing versioned and git crates like this:

```toml
# This will not compile
[dependencies]
gdk = "0.13"
gdk = { git = "https://github.com/gtk-rs/gtk3-rs.git", package = "gdk" }
```

### See Also

 * [glib](https://crates.io/crates/glib)
 * [cairo-rs](https://crates.io/crates/cairo-rs)
 * [gdk-pixbuf](https://crates.io/crates/gdk-pixbuf)
 * [gio](https://crates.io/crates/gio)
 * [pango](https://crates.io/crates/pango)

## License

__gdk__ is available under the MIT License, please refer to it.
