package gliba

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/glib"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/glibi"
)

type settingsSchemaSource struct {
	*glib.SettingsSchemaSource
}

func wrapSettingsSchemaSourceSimple(v *glib.SettingsSchemaSource) *settingsSchemaSource {
	if v == nil {
		return nil
	}
	return &settingsSchemaSource{v}
}

func unwrapSettingsSchemaSource(v glibi.SettingsSchemaSource) *glib.SettingsSchemaSource {
	if v == nil {
		return nil
	}
	return v.(*settingsSchemaSource).SettingsSchemaSource
}

func (v *settingsSchemaSource) Ref() glibi.SettingsSchemaSource {
	return wrapSettingsSchemaSourceSimple(v.SettingsSchemaSource.Ref())
}

func (v *settingsSchemaSource) Unref() {
	v.SettingsSchemaSource.Unref()
}

func (v *settingsSchemaSource) Lookup(v1 string, v2 bool) glibi.SettingsSchema {
	return wrapSettingsSchemaSimple(v.SettingsSchemaSource.Lookup(v1, v2))
}
