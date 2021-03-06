{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/npc">NPC</a> > {{.Npc.CleanName}}</span>
			</div>
			<div class="panel-body">
				<span class="text-center">                    
					<h3> {{.Npc.CleanName}}</h3>                
				</span>
			</div>           
			<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>Spawns In</td><td><a href="/zone/{{.Npc.Zone.ID}}">{{.Npc.Zone.LongName}}</a></td></tr>
					<tr><td>Class</td><td><i title="{{.Npc.Class.Name}}" class="xa {{.Npc.Class.Icon}}"></i> {{.Npc.Class.Name}}</td></tr>
					<tr><td>Race</td><td><i title="{{.Npc.Race.Name}}" class="xa {{.Npc.Race.Icon}}"></i> 
					{{.Npc.Race.Name}}</td></tr>
					<tr><td>Level</td><td>{{.Npc.Level}}</td></tr>
					<tr><td>HP</td><td>{{comma .Npc.Hitpoints}}</td></tr>
					<tr><td>Damage</td><td>{{.Npc.MininumDamage}}-{{.Npc.MaximumDamage}}</td></tr>					
					<tr><td>Experience</td><td><span title="Experience is calculated based on zone experience modifiers and hotzone">{{comma .Npc.Experience}}</span></td></tr>
					<tr><td>Loot ID</td><td>{{.Npc.LootID}}</td></tr>
					{{range $field, $description := .Npc.SpecialAbilities}}                    
						<tr><td>{{$field}}</td><td>{{$description}}</td>
					{{end}}
					</tbody>
				</table>
			</div>
			<div class="panel-footer">
			</div>
		</div>
	</div>
	{{/*if .Map}}
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
			   Map
			</div>
			<div class="panel-body">
				<span class="text-center">
				<svg width="300px" height="300px" version="1.1"
				     xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">
					<image xlink:href="/images/maps/{{.Npc.ZoneName}}.png" x="0" y="0" height="300px" width="300px"/>				
					{{if .Spawns}}
					{{range $key, $value := .Spawns}}
					<circle cx="{{$value.XScaled}}" cy="{{$value.YScaled}}" r="3" fill="red"/>
					{{end}}
					{{end}}
				</svg>
				</span>
				{{if .Spawns}}					
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table table-striped">

						<thead>
						<tr>							
							<th>Name</th>
							<th>Location</th>
						</tr>
						</thead>
						<tbody>
						{{range $key, $value := .Spawns}}
						<tr>
							<td><a href="/spawn/{{$value.SpawngroupID}}">{{$value.Name}}</a></td>
							<td>{{$value.X}}, {{$value.Y}}, {{$value.Z}}</td>							
						</tr>
						{{end}}                
						</tbody>
					
				</table>
				</div>
				{{end}}
			</div>
			{{if .Spawns}}
			<div class="panel-footer">
				{{len .Spawns}} spawn locations
			</div>
			{{end}}
		</div>
	</div>
	{{end*/}}
	{{if .Loot}}
	{{if .Loot.Entrys}}	
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
			   Items
			</div>
			<div class="panel-body">
							
				{{range $entryKey, $entryValue := .Loot.Entrys}}
				<div class="row">
					<div class="col-md-6">
					Multiplier: {{$entryValue.Multiplier}}<br>
					</div>
					<div class="col-md-6">
					Probability: {{$entryValue.Probability}}%<br>
					</div>
				</div>
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table table-striped">

						<thead>
						<tr>
							<th width="20px"><i class="xa xa-sword"></i></th>
							<th>Name</th>	
							<th>%</th>						
						</tr>
						</thead>
						<tbody>
						{{range $key, $value := $entryValue.DropEntrys}}
						<tr>
							<td><span class="item icon-{{$value.Item.Icon}}-sm"></span></td>
							<td><a item={{$value.Item.ID}} href="/item/{{$value.Item.ID}}">{{$value.Item.Name}}</a></td>								
							<td><div data-tooltip="There is a {{$value.Chance}}% chance this item will drop">{{$value.Chance}}%</div></td>
						</tr>
						{{end}}                
						</tbody>
					
				</table>
				</div>
				{{end}}
			</div>
			<div class="panel-footer">
				{{.ItemPage.Total}} items drop from {{.Npc.CleanName}}
			</div>
		</div>

	</div>
	{{end}}
	{{end}}	
	</div>
</div>
</div>